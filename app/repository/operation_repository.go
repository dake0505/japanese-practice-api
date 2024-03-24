package repository

import (
	dao "gin-gonic-api/app/domain/dao/vocabulary"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OperationRepository interface {
	NextOperation(id int) (*dao.N2VocabularySubject, error)
	PreOperation(id int) (*dao.N2VocabularySubject, error)
}

type OperationRepositoryImpl struct {
	db *gorm.DB
}

func (o OperationRepositoryImpl) NextOperation(id int) (*dao.N2VocabularySubject, error) {
	var nextItem dao.N2VocabularySubject
	var err = o.db.Where("id > ?", id).Order("id ASC").First(&nextItem).Error
	if err != nil {
		log.Error("Got an error finding n2 vocabulary", err)
		return nil, err
	}
	return &nextItem, nil
}

func (o OperationRepositoryImpl) PreOperation(id int) (*dao.N2VocabularySubject, error) {
	var preItem dao.N2VocabularySubject
	var err = o.db.Where("id < ?", id).Order("id DESC").First(&preItem).Error
	if err != nil {
		log.Error("Got an error finding n2 vocabulary", err)
		return nil, err
	}
	return &preItem, nil
}

func OperationRepositoryInit(db *gorm.DB) *OperationRepositoryImpl {
	return &OperationRepositoryImpl{
		db: db,
	}
}
