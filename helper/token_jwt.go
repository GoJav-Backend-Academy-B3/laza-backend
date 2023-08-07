package helper

import (
	"github.com/phincon-backend/laza/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte(config.LoadJWTConfig().GetJWTKey())

type claims struct {
	UserId uint64
	Role   bool
	jwt.RegisteredClaims
}

func NewToken(id uint64, role bool) *claims {
	return &claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.LoadJWTConfig().GetTokenExpiry())),
		},
	}

}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecret)
}

func VerifyToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)

	return claims, nil
}
