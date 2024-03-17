package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type OperationController interface {
	Next(c *gin.Context)
}

type OperationControllerImpl struct {
	svc service.OperationService
}

func (o OperationControllerImpl) Next(c *gin.Context) {
	o.svc.NextOperation(c)
}

func OperationControllerInit(operationService service.OperationService) *OperationControllerImpl {
	return &OperationControllerImpl{
		svc: operationService,
	}
}
