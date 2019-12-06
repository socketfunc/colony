package api

import (
	"context"

	"github.com/socketfunc/colony/pkg/auth"
	v1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterAPIServer(srv *grpc.Server, params *ServerParams, opts ...Option) {
	v1beta1.RegisterApiServiceServer(srv, NewServer(params, opts...))
}

type server struct {
	logger *zap.Logger
}

// extends
var _ auth.AuthorizeService = (*server)(nil)
var _ v1beta1.ApiServiceServer = (*server)(nil)

type ServerParams struct{}

func NewServer(params *ServerParams, opts ...Option) v1beta1.ApiServiceServer {
	o := &options{
		logger: zap.NewNop(),
	}
	for i := range opts {
		opts[i](o)
	}
	return &server{
		logger: o.logger,
	}
}

func (s *server) Authorize(ctx context.Context, method string) error {
	return nil
}
