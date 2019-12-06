package api

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	v1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
)

func (s *server) generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	pk, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	if err := pk.Validate(); err != nil {
		return nil, err
	}
	return pk, nil
}

func (s *server) GetPublicKeys(ctx context.Context, req *v1beta1.GetPublicKeysRequest) (*v1beta1.GetPublicKeysResponse, error) {
	return nil, nil
}
