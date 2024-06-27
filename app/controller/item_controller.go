package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	GetItemList(c *gin.Context)
	QueryItemDetail(c *gin.Context)
	CreateQuestionItem(c *gin.Context)
	UpdateQuestionItem(c *gin.Context)
}

type ItemControllerImpl struct {
	svc service.ItemService
}

func (i ItemControllerImpl) GetItemList(c *gin.Context) {
	i.svc.GetItemList(c)
}

func (i ItemControllerImpl) QueryItemDetail(c *gin.Context) {
	idStr := c.Query("questionId")
	userRecord, exists := c.Get("userRecord")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userRecord not found"})
		return
	}
	authUserRecord, ok := userRecord.(*auth.UserRecord)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userRecord type"})
		return
	}
	data := i.svc.QueryQuestionDetail(idStr, authUserRecord.Email)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (i ItemControllerImpl) CreateQuestionItem(c *gin.Context) {
	var body dto.CreateItemRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem := dto.CreateItemRequest{
		QuestionTitle: body.QuestionTitle,
		AnswerId:      body.AnswerId,
	}
	res := i.svc.CreateQuestionItem(newItem)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (i ItemControllerImpl) UpdateQuestionItem(c *gin.Context) {
	var body dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem := dto.UpdateItemRequest{
		QuestionID:    body.QuestionID,
		ID:            body.ID,
		AnswerId:      body.AnswerId,
		QuestionTitle: body.QuestionTitle,
	}
	res := i.svc.UpdateQuestionItem(newItem)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func ItemControllerInit(itemService service.ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{
		svc: itemService,
	}
}
