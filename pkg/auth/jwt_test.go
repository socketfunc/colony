package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	privateKey, publicKey, err := GenerateRSAKeys(2048)
	require.NoError(t, err)

	pr, err := toPrivateKey(privateKey)
	require.NoError(t, err)
	pu, err := toPublicKey(publicKey)
	require.NoError(t, err)

	private := jwt.NewRS256(jwt.RSAPrivateKey(pr))
	public := jwt.NewRS256(jwt.RSAPublicKey(pu))

	pl := &Payload{
		Payload: jwt.Payload{
			Issuer:   "iss",
			Subject:  "sub",
			Audience: jwt.Audience{"https://colony.io"},
			JWTID:    "jti",
		},
		Scopes: []string{"hoge:read", "hoge:write"},
	}
	token, err := New(private).Sign(pl)
	require.NoError(t, err)

	payload, err := New(public).Verify(token)
	require.NoError(t, err)
	assert.Equal(t, pl, payload)
}

func toPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk.Precompute()
	if err := pk.Validate(); err != nil {
		return nil, err
	}
	return pk, nil
}

func toPublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	pk, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("")
	}
	return pk, nil
}
