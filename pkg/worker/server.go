package worker

import (
	"context"

	"github.com/socketfunc/colony/pkg/config"
	"github.com/socketfunc/colony/pkg/runtime"
	"github.com/socketfunc/colony/pkg/storage"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	configpb "github.com/socketfunc/colony/proto/config"
	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterWorkerServer(srv *grpc.Server, params *ServerParams) {
	v1beta1.RegisterWorkerServer(srv, NewServer(params))
}

type server struct {
	apiClient apiv1beta1.ApiServiceClient
	storage   storage.Storage
}

type ServerParams struct {
	ApiClient apiv1beta1.ApiServiceClient
}

func NewServer(params *ServerParams) v1beta1.WorkerServer {
	return &server{
		apiClient: params.ApiClient,
		storage:   storage.New(),
	}
}

func (s *server) Invoke(ctx context.Context, req *v1beta1.InvokeRequest) (*v1beta1.InvokeResponse, error) {
	in := &apiv1beta1.GetConfigRequest{
		Namespace: req.Namespace,
	}
	out, err := s.apiClient.GetConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	cfg := out.Config

	var args []interface{}
	if err := msgpack.Unmarshal(req.Args, &args); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fn := config.Functions(cfg.Spec.Functions).Get(req.Name)
	if fn == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	ret, err := s.invoke(ctx, fn, args)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	retByte, err := msgpack.Marshal(ret)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := &v1beta1.InvokeResponse{
		Result: retByte,
	}
	return resp, nil
}

func (s *server) invoke(ctx context.Context, fn *configpb.Function, args []interface{}) (interface{}, error) {
	file, err := s.storage.Download(fn.Repository)
	if err != nil {
		return nil, err
	}
	r, err := runtime.New(file)
	if err != nil {
		return nil, err
	}
	return r.Execute(fn.Name, args)
}
