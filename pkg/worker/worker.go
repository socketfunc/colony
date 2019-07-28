package worker

import (
	"github.com/socketfunc/colony/pkg/storage"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterWorkerServer(srv *grpc.Server, params *ServerParams, opts ...Option) {
	v1beta1.RegisterWorkerServer(srv, NewServer(params, opts...))
}

type server struct {
	logger    *zap.Logger
	apiClient apiv1beta1.ApiServiceClient
	storage   storage.Storage
}

type ServerParams struct {
	ApiClient apiv1beta1.ApiServiceClient
	Storage   storage.Storage
}

func NewServer(params *ServerParams, opts ...Option) v1beta1.WorkerServer {
	o := &options{
		logger: zap.NewNop(),
	}
	for i := range opts {
		opts[i](o)
	}
	return &server{
		logger:    o.logger,
		apiClient: params.ApiClient,
		storage:   params.Storage,
	}
}

func (s *server) toError(err error) error {
	if err == nil {
		return nil
	}
	return nil
}
