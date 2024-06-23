package router

import (
	"gin-gonic-api/app/middleware"
	"gin-gonic-api/config"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()

	// r.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://japanese-practice-h5.vercel.app"}, // 去掉末尾的斜杠
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.OPTIONS("/*path", func(c *gin.Context) {
		log.Println("Received OPTIONS request")
		c.Header("Access-Control-Allow-Origin", "https://japanese-practice-h5.vercel.app")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(200)
	})

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	auth := api.Group("/auth")
	{
		auth.POST("/login", init.AuthCtrl.Login)
		auth.POST("/register", init.AuthCtrl.Register)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		user := protected.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		n2Vocabulary := protected.Group("/n2Vocabulary")
		n2Vocabulary.GET("/list", init.N2VocabularyCtrl.GetList)
		n2Vocabulary.GET("/:questionId", init.N2VocabularyCtrl.GetQuestionById)

		questionOption := protected.Group("/operation")
		questionOption.GET("/next/:id", init.OperationCtrl.Next)
		questionOption.GET("/pre/:id", init.OperationCtrl.Pre)

		questionType := protected.Group("/type")
		questionType.GET("/list", init.TypeCtrl.GetTypeList)
		questionType.POST("/create", init.TypeCtrl.CreateType)

		item := protected.Group("/item")
		item.GET("/list", init.ItemCtrl.GetItemList)
		item.GET("/detail", init.ItemCtrl.QueryItemDetail)
		item.POST("/create", init.ItemCtrl.CreateQuestionItem)
		item.PUT("/update", init.ItemCtrl.UpdateQuestionItem)

		answer := protected.Group("/answer")
		answer.GET("/list", init.AnswerCtrl.QueryAnswerList)
		answer.POST("/create", init.AnswerCtrl.CreateAnswerItem)

		record := protected.Group("/record")
		record.POST("/create", init.RecordCtrl.CreateRecord)
	}

	return router
}
