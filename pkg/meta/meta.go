package meta

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	namespace = "namespace"
)

func GetNamespace(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	values := md.Get(namespace)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
