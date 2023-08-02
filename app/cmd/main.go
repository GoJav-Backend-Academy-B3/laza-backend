package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/phincon-backend/laza/app/router"
)

func main() {
	apps := router.NewServerGin()

	var address string = "0.0.0.0:8080"
	if APP_PORT := os.Getenv("APP_PORT"); APP_PORT != "" {
		address = "localhost:" + APP_PORT
	}

	apps.Run(address)
}
