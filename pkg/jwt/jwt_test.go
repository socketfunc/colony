package jwt

import (
	"testing"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	secret := jwt.NewHS256([]byte("secret"))
	a := &auth{
		secret: secret,
	}
	pl := &Payload{
		Payload: jwt.Payload{
			Issuer:   "iss",
			Subject:  "sub",
			Audience: jwt.Audience{"https://colony.io"},
			JWTID:    "jti",
		},
		Scopes: []string{"hoge:read", "hoge:write"},
	}
	token, err := a.Sign(pl)
	require.NoError(t, err)

	payload, err := a.Verify(token)
	require.NoError(t, err)
	assert.Equal(t, pl, payload)
}
