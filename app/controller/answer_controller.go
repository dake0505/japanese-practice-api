package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnswerController interface {
	CreateAnswerItem(c *gin.Context)
	QueryAnswerList(c *gin.Context)
}

type AnswerControllerImpl struct {
	answerService service.AnswerService
}

func (a AnswerControllerImpl) QueryAnswerList(c *gin.Context) {
	items := a.answerService.QueryAnswerList()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, items))
}

func (a AnswerControllerImpl) CreateAnswerItem(c *gin.Context) {
	var body dto.CreateAnswerItemDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem := dto.CreateAnswerItemDto{
		AnswerDesc: body.AnswerDesc,
		QuestionID: body.QuestionID,
	}
	res := a.answerService.CreateAnswerItem(newItem)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func AnswerControllerInit(answerService service.AnswerService) *AnswerControllerImpl {
	return &AnswerControllerImpl{
		answerService: answerService,
	}
}
