package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	GetItemList(c *gin.Context)
}

type ItemControllerImpl struct {
	svc service.ItemService
}

func (i ItemControllerImpl) GetItemList(c *gin.Context) {
	i.svc.GetItemList(c)
}

func ItemControllerInit(itemService service.ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{
		svc: itemService,
	}
}
