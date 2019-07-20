package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type gcm struct {
}

func NewGCM() Crypto {
	return &gcm{}
}

func (g *gcm) Type() Type {
	return GCM
}

func (g *gcm) Encrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, toError(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, toError(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, toError(err)
	}
	dst := gcm.Seal(nil, nonce, data, nil)
	dst = append(nonce, dst...)
	return dst, nil
}

func (g *gcm) Decrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, toError(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, toError(err)
	}
	nonce := data[:gcm.NonceSize()]
	dst, err := gcm.Open(nil, nonce, data[gcm.NonceSize():], nil)
	return dst, toError(err)
}
