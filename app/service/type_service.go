package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeService interface {
	GetTypeList(c *gin.Context)
}

type TypeServiceImpl struct {
	typeRepository repository.TypeRepository
}

func (t TypeServiceImpl) GetTypeList(c *gin.Context) {
	defer pkg.PanicHandler(c)
	data, err := t.typeRepository.GetTypeList()
	if err != nil {

	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func TypeServiceInit(typeRepository repository.TypeRepository) *TypeServiceImpl {
	return &TypeServiceImpl{
		typeRepository: typeRepository,
	}
}
