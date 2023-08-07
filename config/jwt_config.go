package config

import (
	"os"
	"strconv"
	"time"
)

type config struct {
	tokenExpiry time.Duration
	jwtKey      string
}

var jwtConfig = new(config)

func init() {
	intExpiry, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
	jwtConfig.jwtKey = os.Getenv("JWT_KEYS")
	jwtConfig.tokenExpiry = time.Duration(intExpiry) * time.Minute
}

func LoadJWTConfig() *config {
	return jwtConfig
}
func (jc *config) GetTokenExpiry() time.Duration {
	return jwtConfig.tokenExpiry
}

func (jc *config) GetJWTKey() string {
	return jwtConfig.jwtKey
}
