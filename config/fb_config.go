package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"log"
	"os"
)

type Config struct {
	LoginConfig oauth2.Config
}

var FBConfig Config

const OauthFacebookUrlAPI = "https://graph.facebook.com/v17.0/me?fields=id,name,first_name,email,picture&access_token="

func init() {
	LoadFBAuthConfig()
}
func LoadFBAuthConfig() {
	log.Println(os.Getenv("FB_CLIENT_SECRET"))
	FBConfig.LoginConfig = oauth2.Config{
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Endpoint:     facebook.Endpoint,
		RedirectURL:  os.Getenv("FB_AUTH_REDIRECT_URL"),
		Scopes: []string{
			"public_profile",
			"email",
		},
	}
}
