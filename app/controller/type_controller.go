package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type TypeContoller interface {
	GetTypeList(c *gin.Context)
}

type TypeContollerImpl struct {
	svc service.TypeService
}

func (t TypeContollerImpl) GetTypeList(c *gin.Context) {
	t.svc.GetTypeList(c)
}

func TypeContollerInit(typeService service.TypeService) *TypeContollerImpl {
	return &TypeContollerImpl{
		svc: typeService,
	}
}
