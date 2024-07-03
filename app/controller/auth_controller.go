package controller

import (
	"gin-gonic-api/app/constant"
	dao "gin-gonic-api/app/domain/dao/auth"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	SendMail(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (a AuthControllerImpl) Login(c *gin.Context) {
	var request dao.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Happened error when mapping request from FE. Error: %v", err)
		c.Error(err)
		return
	}
	customToken, err := a.svc.Login(c, request)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{"token": customToken}))
}

func (a AuthControllerImpl) Register(c *gin.Context) {
	var request dao.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	res, err := a.svc.Register(c, request)
	if err != nil {
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (a AuthControllerImpl) SendMail(c *gin.Context) {
	var request dao.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	res, err := a.svc.SendMail(c, request)
	if err != nil {
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
