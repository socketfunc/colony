package auth

//go:generate mkdir -p mock
//go:generate mockgen -source=jwt.go -package=mock -destination=mock/jwt.go
import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	errors "golang.org/x/xerrors"
)

type Payload struct {
	jwt.Payload
	Scopes Scopes `json:"scopes,omitempty"`
}

type Scopes []string

func (s Scopes) Contains(scope string) bool {
	for i := range s {
		if s[i] == scope {
			return true
		}
	}
	return false
}

type JWTs map[string]JWT

func (j JWTs) Verify(namespace, token string) (*Payload, error) {
	jwt, ok := j[namespace]
	if !ok {
		return nil, errors.New("")
	}
	return jwt.Verify(token)
}

type JWT interface {
	Sign(payload *Payload) (string, error)
	Verify(token string) (*Payload, error)
}

type auth struct {
	secret jwt.Algorithm
	now    time.Time
}

func New(secret jwt.Algorithm) JWT {
	return &auth{
		secret: secret,
		now:    time.Now(),
	}
}

func (a *auth) Sign(payload *Payload) (string, error) {
	expiredAt := a.now.Add(time.Hour).Truncate(time.Second)
	payload.ExpirationTime = &jwt.Time{Time: expiredAt}
	token, err := jwt.Sign(payload, a.secret)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (a *auth) Verify(token string) (*Payload, error) {
	pl := new(Payload)
	_, err := jwt.Verify([]byte(token), a.secret, pl)
	if pl.ExpirationTime.Before(a.now) {
		return nil, errors.New("expired")
	}
	return pl, err
}
