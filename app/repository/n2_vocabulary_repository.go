package repository

import (
	"gin-gonic-api/app/domain/dao"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type N2VocabularyRepository interface {
	GetList() ([]dao.N2Vocabulary, error)
}

type N2VocabularyRepositoryImpl struct {
	db *gorm.DB
}

func (i N2VocabularyRepositoryImpl) GetList() ([]dao.N2Vocabulary, error) {
	var iterms []dao.N2Vocabulary

	var err = i.db.Find(&iterms).Error

	if err != nil {
		log.Error("Got an error finding n2 vocabulary", err)
		return nil, err
	}
	return iterms, nil
}

func N2VocabularyRepositoryInit(db *gorm.DB) *N2VocabularyRepositoryImpl {
	return &N2VocabularyRepositoryImpl{
		db: db,
	}
}
