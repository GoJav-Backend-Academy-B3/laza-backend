package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/phincon-backend/laza/cmd/server/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	apps := router.NewServerGin()
	fmt.Println("listening port 8080")
	http.ListenAndServe("localhost:8080", apps)
}
