package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()
	r.Use(cors.Default())
	init := config.Init(&gin.Context{})
	app := router.Init(init)

	app.Run(":" + port)
}
