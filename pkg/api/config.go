package api

import (
	"context"

	v1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
)

func (s *server) GetConfig(ctx context.Context, req *v1beta1.GetConfigRequest) (*v1beta1.GetConfigResponse, error) {
	return nil, nil
}

func (s *server) SaveConfig(ctx context.Context, req *v1beta1.SaveConfigRequest) (*v1beta1.SaveConfigResponse, error) {
	return nil, nil
}
