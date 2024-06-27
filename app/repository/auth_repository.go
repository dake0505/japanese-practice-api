package repository

import (
	dao "gin-gonic-api/app/domain/dao/auth"

	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindAuthByEmail(email string) (dao.Auth, error)
	CreateUser(auth *dao.Auth) (dao.Auth, error)
}

type AuthRepositoryImpl struct {
	db       *gorm.DB
	fireAuth *firebase.App
}

func (a AuthRepositoryImpl) FindAuthByEmail(email string) (dao.Auth, error) {
	auth := dao.Auth{
		Email: email,
	}
	err := a.db.First(&auth).Error
	log.Error("Got and error when find user by email. Error: ", err)
	if err != nil {
		log.Error("Got and error when find user by email. Error: ", err)
		return dao.Auth{}, err
	}
	return auth, nil
}

func (a AuthRepositoryImpl) CreateUser(auth *dao.Auth) (dao.Auth, error) {
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
