package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
}

func (a AuthServiceImpl) Login(c *gin.Context) {
	email := (c.Param("email"))
	// Get the user from the database
	data, err := a.authRepository.FindAuthByEmail(email)
	if err != nil {
		log.Printf("failed to get user by email from database: %v", err)
	}

	token, err := a.authRepository.FindToken(data.Email)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{"token": token}))
}

func (a AuthServiceImpl) Register(c *gin.Context) {
	email := (c.Param("email"))
	password := (c.Param("password"))
	// Check if the user with the email already exists

	data, err := a.authRepository.FindAuthByEmail(email)

	if data.Email != "" {
		errors.New("user with email already exists")
	}

	// Generate a UUID for the new user
	uid := uuid.New().String()

	// Generate a hash of the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		errors.New("internal server error")
	}

	var newUser dao.Auth
	newUser.ID = uid
	newUser.Email = email
	newUser.Password = string(hashedPassword)

	// Create a new user in the database
	if _, err := a.authRepository.Save(&newUser); err != nil {
		log.Printf("failed to insert user into database: %v", err)
		errors.New("internal server error")
	}

	customToken, err := a.authRepository.CreateToken(uid)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{"token": customToken}))
}

func AuthServiceInit(authRepository repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
	}
}
