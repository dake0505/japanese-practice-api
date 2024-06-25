package service

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	dao "gin-gonic-api/app/domain/dao/auth"
	"gin-gonic-api/app/repository"
)

type AuthService interface {
	Login(c *gin.Context, userInfo dao.Auth) (string, error)
	Register(c *gin.Context, userInfo dao.Auth) (string, error)
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	fireAuth       *firebase.App
}

func (a AuthServiceImpl) Login(c *gin.Context, userInfo dao.Auth) (string, error) {
	email := userInfo.Email
	password := userInfo.Password
	data, err := a.authRepository.FindAuthByEmail(email)
	log.Printf("failed to get user by email from database: %v", err)
	if err != nil {
		log.Printf("failed to get user by email from database: %v", err)
		return "failed", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		log.Printf("password does not match: %v", err)
		return "failed", err
	}
	client, err := a.fireAuth.Auth(context.Background())
	if err != nil {
		log.Printf("failed to get user by email from Firebase: %v", err)
		return "failed", err
	}
	u, err := client.GetUserByEmail(c, email)
	customToken, err := client.CustomToken(c.Request.Context(), u.UID)
	if err != nil {
		log.Printf("failed to create custom token: %v", err)
		return "failed", err
	}
	return customToken, nil
}

func (a AuthServiceImpl) Register(c *gin.Context, userInfo dao.Auth) (string, error) {
	email := userInfo.Email
	password := userInfo.Password
	data, err := a.authRepository.FindAuthByEmail(email)
	if data.Email != "" {
		return "failed", errors.New("user with email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return "failed", errors.New("internal server error")
	}

	var newUser dao.Auth
	newUser.Email = email
	newUser.Password = string(hashedPassword)
	if _, err := a.authRepository.CreateUser(&newUser); err != nil {
		log.Printf("failed to insert user into database: %v", err)
		return "failed", errors.New("internal server error")
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
		return "failed", err
	}
	u, err := client.CreateUser(c, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
		return "failed", err
	}
	log.Printf("Successfully created user: %v\n", u)
	return "ok", err
}

func AuthServiceInit(authRepository repository.AuthRepository, fireAuth *firebase.App) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
		fireAuth:       fireAuth,
	}
}
