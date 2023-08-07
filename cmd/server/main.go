package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/phincon-backend/laza/cmd/server/router"
	"github.com/phincon-backend/laza/config"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	apps := router.NewServerGin()
	fmt.Println("listening port 8080")
	http.ListenAndServe("localhost:8080", apps)

	config.LoadConfig()
	apps := router.NewServerGin()
	var address string = "0.0.0.0:8080"
	if APP_PORT := os.Getenv("APP_PORT"); APP_PORT != "" {
		address = "localhost:" + APP_PORT
	}

	apps.Run(address)

}
