package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	GetItemList(c *gin.Context)
	CreateQuestionItem(c *gin.Context)
}

type ItemControllerImpl struct {
	svc service.ItemService
}

func (i ItemControllerImpl) GetItemList(c *gin.Context) {
	i.svc.GetItemList(c)
}

func (i ItemControllerImpl) CreateQuestionItem(c *gin.Context) {
	var body dto.CreateItemRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理请求数据
	newItem := dto.CreateItemRequest{
		QuestionTitle: body.QuestionTitle,
		AnswerId:      body.AnswerId,
	}
	res := i.svc.CreateQuestionItem(newItem)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func ItemControllerInit(itemService service.ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{
		svc: itemService,
	}
}
