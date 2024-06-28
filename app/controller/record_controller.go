package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/firebase"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type RecordController interface {
	QueryRecordList(c *gin.Context)
	CreateRecord(c *gin.Context)
	UpdateFavorite(c *gin.Context)
}

type RecordControllerImpl struct {
	recordService service.RecordService
}

func (r RecordControllerImpl) QueryRecordList(c *gin.Context) {
	userRecord, exists := c.Get("userRecord")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userRecord not found"})
		return
	}
	authUserRecord, ok := userRecord.(*auth.UserRecord)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userRecord type"})
		return
	}
	params := dto.CreateRecordDto{
		CreatedBy:  authUserRecord.Email,
		RecordType: c.Query("recordType"),
	}
	res, err := r.recordService.QueryRecordList(params)
	if err != nil {
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
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

func (r RecordControllerImpl) UpdateFavorite(c *gin.Context) {
	userRecord, exists := c.Get("userRecord")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userRecord not found"})
		return
	}
	authUserRecord, ok := userRecord.(*auth.UserRecord)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userRecord type"})
		return
	}
	log.Printf("authUserRecord, %v", authUserRecord.Email)
	var body dto.UpdateFavoriteDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := dto.UpdateFavoriteDto{
		QuestionId: body.QuestionId,
		CreatedBy:  authUserRecord.Email,
	}
	res, err := r.recordService.UpdateRecord(data, "favorite")
	if err != nil {
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func RecordControllerInit(recordService service.RecordService) *RecordControllerImpl {
	return &RecordControllerImpl{
		recordService: recordService,
	}
}
