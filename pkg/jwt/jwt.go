package jwt

//go:generate mkdir -p mock
//go:generate mockgen -source=jwt.go -package=mock -destination=mock/jwt.go
import (
	"github.com/gbrlsnchs/jwt/v3"
)

type Payload struct {
	jwt.Payload
	Scopes []string
}

type JWT interface {
	Sign(payload *Payload) (string, error)
	Verify(token string) (*Payload, error)
}

type auth struct {
	secret jwt.Algorithm
}

func (a *auth) Sign(payload *Payload) (string, error) {
	token, err := jwt.Sign(payload, a.secret)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (a *auth) Verify(token string) (*Payload, error) {
	pl := new(Payload)
	_, err := jwt.Verify([]byte(token), a.secret, pl)
	return pl, err
}
