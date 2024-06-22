package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"os"
	"time"

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

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://japanese-practice-h5-ne6uh5zk4-dake0505s-projects.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	init := config.Init(&gin.Context{})
	app := router.Init(init)

	app.Run(":" + port)
}
