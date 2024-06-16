package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_type"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TypeRepository interface {
	GetTypeList() ([]dao.QuestionType, error)
}

type TypeRepositoryImpl struct {
	db *gorm.DB
}

func (t TypeRepositoryImpl) GetTypeList() ([]dao.QuestionType, error) {
	var types []dao.QuestionType
	var err = t.db.Find(&types).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}
	log.Println(types, "lllllllllllllllll")
	return types, nil
}

func TypeRepositoryInit(db *gorm.DB) *TypeRepositoryImpl {
	return &TypeRepositoryImpl{
		db: db,
	}
}
