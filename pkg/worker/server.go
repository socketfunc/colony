package worker

import (
	"context"

	"github.com/socketfunc/colony/pkg/config"
	"github.com/socketfunc/colony/pkg/runtime"
	"github.com/socketfunc/colony/pkg/storage"
	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterWorkerServer(srv *grpc.Server) {
	s := NewServer()
	// TODO: sample data
	ss := s.(*server)
	cfg := &config.Config{
		Version: "v1beta1",
		Metadata: map[string]string{
			"metadata": "default",
		},
		Spec: &config.Spec{
			Functions: []*config.Function{
				{
					Name:       "add",
					Repository: "file:///tmp/sample.wasm",
					Resources:  &config.Resources{},
				},
				{
					Name:       "hello",
					Repository: "file:///tmp/sample.wasm",
					Resources:  &config.Resources{},
				},
			},
		},
	}
	ss.configStore.Save("default", cfg)
	v1beta1.RegisterWorkerServer(srv, s)
}

type server struct {
	storage     storage.Storage
	configStore *config.Store
}

func NewServer() v1beta1.WorkerServer {
	return &server{
		storage:     storage.New(),
		configStore: &config.Store{},
	}
}

func (s *server) Invoke(ctx context.Context, req *v1beta1.InvokeRequest) (*v1beta1.InvokeResponse, error) {
	cfg := s.configStore.Get(req.Namespace)
	if cfg == nil {
		return nil, status.Error(codes.NotFound, "")
	}

	var args []interface{}
	if err := msgpack.Unmarshal(req.Args, &args); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fn := cfg.Spec.GetFunction(req.Name)
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

func (s *server) invoke(ctx context.Context, fn *config.Function, args []interface{}) (interface{}, error) {
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
