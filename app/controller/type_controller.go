package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeController interface {
	GetTypeList(c *gin.Context)
	CreateType(c *gin.Context)
}

type TypeControllerImpl struct {
	svc service.TypeService
}

func (t TypeControllerImpl) GetTypeList(c *gin.Context) {
	t.svc.GetTypeList(c)
}

func (t TypeControllerImpl) CreateType(c *gin.Context) {
	var body dto.CreateTypeDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem := dto.CreateTypeDto{
		TypeName: body.TypeName,
	}
	res := t.svc.CreateType(newItem)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func TypeControllerInit(typeService service.TypeService) *TypeControllerImpl {
	return &TypeControllerImpl{
		svc: typeService,
	}
}
