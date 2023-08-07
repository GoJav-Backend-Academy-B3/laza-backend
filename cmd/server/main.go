package main

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitterv2"
	"github.com/phincon-backend/laza/helper"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/phincon-backend/laza/cmd/server/router"
	"github.com/phincon-backend/laza/config"
)

func main() {
	config.LoadConfig()

	gothic.Store = helper.GetStore()
	goth.UseProviders(
		twitterv2.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://127.0.0.1:8080/auth/twitter/callback?provider=twitterv2"),
	)
	apps := router.NewServerGin()
	var address string = "0.0.0.0:8080"
	if APP_PORT := os.Getenv("APP_PORT"); APP_PORT != "" {
		address = "localhost:" + APP_PORT
	}

	apps.Run(address)
}
