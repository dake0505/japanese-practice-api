package repository

import (
	"context"
	"fmt"
	dao "gin-gonic-api/app/domain/dao/auth"

	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(auth *dao.Auth) (dao.Auth, error)
	Register(auth *dao.Auth) (dao.Auth, error)
	FindAuthByEmail(email string) (dao.Auth, error)
	FindToken(email string) (string, error)
	CreateToken(uid string) (string, error)
	Save(auth *dao.Auth) (dao.Auth, error)
}

type AuthRepositoryImpl struct {
	db       *gorm.DB
	fireAuth *firebase.App
}

func (a AuthRepositoryImpl) Login(auth *dao.Auth) (dao.Auth, error) {
	var err = a.db.Save(auth).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.Auth{}, err
	}
	return *auth, nil
}

func (a AuthRepositoryImpl) Register(auth *dao.Auth) (dao.Auth, error) {
	var err = a.db.Save(auth).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.Auth{}, err
	}
	return *auth, nil
}

func (a AuthRepositoryImpl) FindAuthByEmail(email string) (dao.Auth, error) {
	auth := dao.Auth{
		Email: email,
	}
	err := a.db.First(&auth).Error
	if err != nil {
		log.Error("Got and error when find user by email. Error: ", err)
		return dao.Auth{}, err
	}
	return auth, nil
}

func (a AuthRepositoryImpl) FindToken(email string) (string, error) {
	client, err := a.fireAuth.Auth(context.Background())
	token, err := client.CustomToken(context.Background(), email)
	if err != nil {
		log.Printf("failed to generate custom token: %v", err)
	}
	return token, err
}

func (a *AuthRepositoryImpl) CreateToken(uid string) (string, error) {
	client, err := a.fireAuth.Auth(context.Background())
	if err != nil {
		log.Errorf("Failed to get Firebase auth client: %v", err)
		return "", fmt.Errorf("failed to get Firebase auth client: %v", err)
	}

	customToken, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		log.Errorf("Failed to create custom token for user: %v", err)
		return "", fmt.Errorf("failed to create custom token for user: %v", err)
	}
	return customToken, nil
}

func (a AuthRepositoryImpl) Save(auth *dao.Auth) (dao.Auth, error) {
	var err = a.db.Save(auth).Error
	if err != nil {
		log.Error("Got an error when save auth. Error: ", err)
		return dao.Auth{}, err
	}
	return *auth, nil
}

func AuthRepositoryInit(db *gorm.DB, fireAuth *firebase.App) *AuthRepositoryImpl {
	//db.AutoMigrate(&dao.User{})
	return &AuthRepositoryImpl{
		db:       db,
		fireAuth: fireAuth,
	}
}
