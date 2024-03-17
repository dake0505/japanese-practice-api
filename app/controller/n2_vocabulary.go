package controller

import (
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type N2VocabularyController interface {
	GetList(c *gin.Context)
	GetQuestionById(c *gin.Context)
}

type N2VocabularyControllerImpl struct {
	svc service.N2VocabularyService
}

func (n N2VocabularyControllerImpl) GetList(c *gin.Context) {
	p := pkg.PaginatorHandler(c)
	n.svc.GetList(c, p)
}

func (n N2VocabularyControllerImpl) GetQuestionById(c *gin.Context) {
	p := pkg.PaginatorHandler(c)
	questionId := cast.ToInt(c.Param("questionId"))
	n.svc.GetQuestionById(c, p, questionId)
}

func N2VocabularyControllerInit(n2VocabularyService service.N2VocabularyService) *N2VocabularyControllerImpl {
	return &N2VocabularyControllerImpl{svc: n2VocabularyService}
}
