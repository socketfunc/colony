package crypto

import "golang.org/x/xerrors"

type Type int

const (
	GCM Type = iota + 1
)

type Crypto interface {
	Type() Type
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
}

func toError(err error) error {
	if err == nil {
		return nil
	}
	return xerrors.Errorf("crypto: %w", err)
}
