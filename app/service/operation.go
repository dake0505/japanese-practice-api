package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type OperationService interface {
	NextOperation(c *gin.Context)
}

type OperationServiceImpl struct {
	operationRepository repository.OperationRepository
}

func (o OperationServiceImpl) NextOperation(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute program get all n2 vocabulary list")

	data := o.operationRepository.NextOperation()

	// if err != nil {
	// 	log.Error("Happened Error when find all n2 vocabulary list. Error: ", err)
	// 	pkg.PanicException((constant.UnknownError))
	// }

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func OperationServiceInit(operationRepository repository.OperationRepository) *OperationServiceImpl {
	return &OperationServiceImpl{
		operationRepository: operationRepository,
	}
}
