package router

import (
	"gin-gonic-api/app/middleware"
	"gin-gonic-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
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
