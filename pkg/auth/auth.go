package auth

import (
	"context"
	"strings"

	errors "golang.org/x/xerrors"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationKey = "authorization"
)

type Roles map[string]string

func (r Roles) Get(method string) string {
	return r[method]
}

func GetToken(ctx context.Context, scheme string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("")
	}
	v := md.Get(AuthorizationKey)
	if len(v) == 0 {
		return "", errors.New("")
	}
	auth := strings.SplitN(v[0], " ", 2)
	if len(auth) == 0 {
		return "", errors.New("")
	}
	if strings.ToLower(auth[0]) != scheme {
		return "", errors.New("")
	}
	return auth[1], nil
}
