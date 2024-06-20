package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/firebase"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordController interface {
	CreateRecord(c *gin.Context)
}

type RecordControllerImpl struct {
	recordService service.RecordService
}

func (r RecordControllerImpl) CreateRecord(c *gin.Context) {
	userAuthID, exists := c.Get("userAuthID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userAuthID not found"})
		return
	}
	app := firebase.InitFirebase()
	client, err := app.Auth(c.Request.Context())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	userRecord, err := client.GetUser(c.Request.Context(), userAuthID.(string))
	var body dto.CreateRecordDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newRecord := dto.CreateRecordDto{
		QuestionId:   body.QuestionId,
		AnswerId:     body.AnswerId,
		UserAnswerId: body.UserAnswerId,
		RecordType:   body.RecordType,
		CreatedBy:    userRecord.Email,
	}
	res := r.recordService.CreateRecord(newRecord)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func RecordControllerInit(recordService service.RecordService) *RecordControllerImpl {
	return &RecordControllerImpl{
		recordService: recordService,
	}
}
