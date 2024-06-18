package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init(&gin.Context{})
	app := router.Init(init)

	app.Run(":" + port)
}
