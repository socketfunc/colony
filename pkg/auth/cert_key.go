package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKeys(size int) (privateKey []byte, publicKey []byte, err error) {
	key, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, nil, err
	}
	if err := key.Validate(); err != nil {
		return nil, nil, err
	}

	privateKey = encodePrivateKeyToPEM(key)
	publicKey, err = encodePublicKeyToPEM(&key.PublicKey)
	return
}

func encodePrivateKeyToPEM(key *rsa.PrivateKey) []byte {
	pk := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.EncodeToMemory(pk)
}

func encodePublicKeyToPEM(key *rsa.PublicKey) ([]byte, error) {
	bytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}

	pk := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}
	return pem.EncodeToMemory(pk), nil
}
