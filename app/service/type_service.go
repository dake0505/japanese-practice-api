package service

import (
	"gin-gonic-api/app/constant"
	dao "gin-gonic-api/app/domain/dao/question_type"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeService interface {
	GetTypeList(c *gin.Context)
	CreateType(item dto.CreateTypeDto) dao.QuestionType
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

func (t TypeServiceImpl) CreateType(item dto.CreateTypeDto) dao.QuestionType {
	data, err := t.typeRepository.CreateType(&item)
	if err != nil {
	}
	return data
}

func TypeServiceInit(typeRepository repository.TypeRepository) *TypeServiceImpl {
	return &TypeServiceImpl{
		typeRepository: typeRepository,
	}
}
