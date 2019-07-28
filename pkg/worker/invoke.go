package worker

import (
	"context"

	"github.com/socketfunc/colony/pkg/config"
	"github.com/socketfunc/colony/pkg/runtime"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	configpb "github.com/socketfunc/colony/proto/config"
	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Invoke(ctx context.Context, req *v1beta1.InvokeRequest) (*v1beta1.InvokeResponse, error) {
	cfg, err := s.getConfig(ctx, req.Namespace)
	if err != nil {
		return nil, err
	}

	args, err := s.decode(req.Args)
	if err != nil {
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

	retByte, err := s.encode(ret)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := &v1beta1.InvokeResponse{
		Result: retByte,
	}
	return resp, nil
}

func (s *server) getConfig(ctx context.Context, namespace string) (*configpb.Config, error) {
	in := &apiv1beta1.GetConfigRequest{
		Namespace: namespace,
	}
	out, err := s.apiClient.GetConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out.Config, nil
}

func (s *server) decode(data []byte) ([]interface{}, error) {
	var ret []interface{}
	if err := msgpack.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *server) encode(data interface{}) ([]byte, error) {
	return msgpack.Marshal(data)
}

func (s *server) invoke(ctx context.Context, fn *configpb.Function, args []interface{}) (interface{}, error) {
	file, err := s.storage.Download(ctx, fn.Repository)
	if err != nil {
		return nil, err
	}
	r, err := runtime.New(file)
	if err != nil {
		return nil, err
	}
	return r.Execute(fn.Name, args)
}
