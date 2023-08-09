package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"
	"github.com/phincon-backend/laza/cmd/server/router"
	"github.com/phincon-backend/laza/config"
)

func main() {
	config.LoadConfig()

	godotenv.Load("/Users/juandaantoniuspakpahan/phincon/week5/project/laza-backend/.env")

	// gothic.Store = helper.GetStore()
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), os.Getenv("TWITTER_REDIRECT_URI")),
	)

	apps := router.NewServerGin()
	var address string = "0.0.0.0:8080"
	if APP_PORT := os.Getenv("APP_PORT"); APP_PORT != "" {
		address = "localhost:" + APP_PORT
	}

	apps.Run(address)
}
