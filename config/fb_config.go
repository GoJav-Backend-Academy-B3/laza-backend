package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"os"
)

type Config struct {
	LoginConfig oauth2.Config
}

var FBConfig Config

const OauthFacebookUrlAPI = "https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token="

func init() {
	LoadFBAuthConfig()
}
func LoadFBAuthConfig() {
	FBConfig.LoginConfig = oauth2.Config{
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/facebook/callback",
		Scopes: []string{
			"email",
			"public_profile",
		},
	}
}
