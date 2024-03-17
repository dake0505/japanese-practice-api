package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type OperationController interface {
	Next(c *gin.Context)
	Pre(c *gin.Context)
}

type OperationControllerImpl struct {
	svc service.OperationService
}

func (o OperationControllerImpl) Next(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	o.svc.NextOperation(c, id)
}

func (o OperationControllerImpl) Pre(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	o.svc.PreOperation(c, id)
}

func OperationControllerInit(operationService service.OperationService) *OperationControllerImpl {
	return &OperationControllerImpl{
		svc: operationService,
	}
}
