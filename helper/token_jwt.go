package helper

import (
	"time"

	"github.com/phincon-backend/laza/config"

	"github.com/golang-jwt/jwt/v4"
)

var mySecretToken = []byte(config.LoadJWTConfig().GetJWTTokenKey())
var mySecretRefresh = []byte(config.LoadJWTConfig().GetJWTRefreshKey())

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

func NewRefresh(id uint64, role bool) *claims {
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

	return tokens.SignedString(mySecretToken)
}

func (c *claims) CreateRefresh() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecretRefresh)
}

func VerifyToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretToken), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)

	return claims, nil
}

func VerifyRefresh(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretRefresh), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)

	return claims, nil
}
