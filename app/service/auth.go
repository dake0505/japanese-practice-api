package service

import (
	"context"
	"errors"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"gin-gonic-api/app/constant"
	dao "gin-gonic-api/app/domain/dao/auth"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
)

type AuthService interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	fireAuth       *firebase.App
}

func (a AuthServiceImpl) Login(c *gin.Context) {
	var request dao.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	email := request.Email
	password := request.Password
	data, err := a.authRepository.FindAuthByEmail(email)
	if err != nil {
		log.Printf("failed to get user by email from database: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		log.Printf("password does not match: %v", err)
	}
	client, err := a.fireAuth.Auth(context.Background())
	u, err := client.GetUserByEmail(c, email)
	customToken, err := client.CustomToken(c.Request.Context(), u.UID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{"token": customToken}))
}

func (a AuthServiceImpl) Register(c *gin.Context) {
	var request dao.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	email := request.Email
	password := request.Password
	data, err := a.authRepository.FindAuthByEmail(email)
	if data.Email != "" {
		errors.New("user with email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		errors.New("internal server error")
	}

	var newUser dao.Auth
	newUser.Email = email
	newUser.Password = string(hashedPassword)
	if _, err := a.authRepository.CreateUser(&newUser); err != nil {
		log.Printf("failed to insert user into database: %v", err)
		errors.New("internal server error")
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(string(hashedPassword)).
		DisplayName("John").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	client, err := a.fireAuth.Auth(context.Background())
	if err != nil {
		log.Errorf("Failed to get Firebase auth client: %v", err)
	}
	u, err := client.CreateUser(c, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func AuthServiceInit(authRepository repository.AuthRepository, fireAuth *firebase.App) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		fireAuth:       fireAuth,
	}
}
