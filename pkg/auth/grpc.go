package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorizeService interface {
	Authorize(ctx context.Context, method string) error
}

func AuthorizeInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		srv, ok := info.Server.(AuthorizeService)
		if !ok {
			return nil, status.Error(codes.Internal, "")
		}
		err := srv.Authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
