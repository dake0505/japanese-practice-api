package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"log"
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
	// r.Use(cors.Default())
	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://japanese-practice-h5-ne6uh5zk4-dake0505s-projects.vercel.app/"}, // 修改为你允许的前端域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 处理预检请求
	r.OPTIONS("/*path", func(c *gin.Context) {
		log.Println("Received OPTIONS request")
		c.Header("Access-Control-Allow-Origin", "https://japanese-practice-h5-ne6uh5zk4-dake0505s-projects.vercel.app")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(204)
	})

	init := config.Init(&gin.Context{})
	app := router.Init(init)

	app.Run(":" + port)
}
