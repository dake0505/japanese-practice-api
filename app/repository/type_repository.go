package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_type"
	"gin-gonic-api/app/domain/dto"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TypeRepository interface {
	GetTypeList() ([]dao.QuestionType, error)
	CreateType(item *dto.CreateTypeDto) (dao.QuestionType, error)
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
	return types, nil
}

func (t TypeRepositoryImpl) CreateType(item *dto.CreateTypeDto) (dao.QuestionType, error) {
	questionType := dao.QuestionType{
		TypeName: item.TypeName,
	}
	err := t.db.Create(&questionType).Error
	if err != nil {
		return dao.QuestionType{}, err
	}
	return questionType, nil
}

func TypeRepositoryInit(db *gorm.DB) *TypeRepositoryImpl {
	return &TypeRepositoryImpl{
		db: db,
	}
}
