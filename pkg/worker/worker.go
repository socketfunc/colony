package worker

import (
	"context"

	"github.com/socketfunc/colony/pkg/auth"
	"github.com/socketfunc/colony/pkg/meta"
	"github.com/socketfunc/colony/pkg/storage"
	apiv1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterWorkerServer(srv *grpc.Server, params *ServerParams, opts ...Option) {
	v1beta1.RegisterWorkerServer(srv, NewServer(params, opts...))
}

type server struct {
	logger    *zap.Logger
	jwts      auth.JWTs
	roles     auth.Roles
	apiClient apiv1beta1.ApiServiceClient
	storage   storage.Storage
}

var _ auth.AuthorizeService = (*server)(nil)
var _ v1beta1.WorkerServer = (*server)(nil)

type ServerParams struct {
	JWTs      auth.JWTs
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
		jwts:      params.JWTs,
		apiClient: params.ApiClient,
		storage:   params.Storage,
		roles: auth.Roles{
			"/colony.worker.v1beta1.Worker/Invoke": "worker.invoke",
		},
	}
}

func (s *server) Authorize(ctx context.Context, method string) error {
	const scheme = "bearer"
	namespace := meta.GetNamespace(ctx)
	token, err := auth.GetToken(ctx, scheme)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}
	payload, err := s.jwts.Verify(namespace, token)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}
	scope := s.roles.Get(method)
	if !payload.Scopes.Contains(scope) {
		return status.Error(codes.PermissionDenied, "")
	}
	return nil
}
