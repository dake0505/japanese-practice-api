package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type N2VocabularyService interface {
	GetList(c *gin.Context, p *pkg.Paginator)
	GetQuestionById(c *gin.Context, p *pkg.Paginator)
}

type N2VocabularyServiceImpl struct {
	n2VocabularyRepository repository.N2VocabularyRepository
}

func (n N2VocabularyServiceImpl) GetList(c *gin.Context, p *pkg.Paginator) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get all n2 vocabulary list")

	data, err := n.n2VocabularyRepository.GetList(p)

	if err != nil {
		log.Error("Happened Error when find all n2 vocabulary list. Error: ", err)
		pkg.PanicException((constant.UnknownError))
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (n N2VocabularyServiceImpl) GetQuestionById(c *gin.Context, p *pkg.Paginator) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get all n2 vocabulary list")

	data, err := n.n2VocabularyRepository.GetQuestionById(p)

	if err != nil {
		log.Error("Happened Error when find all n2 vocabulary list. Error: ", err)
		pkg.PanicException((constant.UnknownError))
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func N2VocabularyServiceInit(n2VocabularyRepository repository.N2VocabularyRepository) *N2VocabularyServiceImpl {
	return &N2VocabularyServiceImpl{
		n2VocabularyRepository: n2VocabularyRepository,
	}
}
