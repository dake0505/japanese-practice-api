package router

import (
	"gin-gonic-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		n2Vocabulary := api.Group("/n2Vocabulary")
		n2Vocabulary.GET("/list", init.N2VocabularyCtrl.GetList)
		n2Vocabulary.GET("/:questionId", init.N2VocabularyCtrl.GetQuestionById)

		questionOption := api.Group("/operation")
		questionOption.GET("/next/:id", init.OperationCtrl.Next)
		questionOption.GET("/pre/:id", init.OperationCtrl.Pre)
	}

	return router
}
