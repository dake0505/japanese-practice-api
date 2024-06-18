package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type TypeController interface {
	GetTypeList(c *gin.Context)
}

type TypeControllerImpl struct {
	svc service.TypeService
}

func (t TypeControllerImpl) GetTypeList(c *gin.Context) {
	t.svc.GetTypeList(c)
}

func TypeControllerInit(typeService service.TypeService) *TypeControllerImpl {
	return &TypeControllerImpl{
		svc: typeService,
	}
}
