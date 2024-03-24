package repository

import (
	dao "gin-gonic-api/app/domain/dao/vocabulary"
	"gin-gonic-api/app/pkg"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type N2VocabularyRepository interface {
	GetList(p *pkg.Paginator) ([]dao.N2VocabularySubject, error)
	GetQuestionById(p *pkg.Paginator, questionId int) ([]dao.N2VocabularySubjectOption, error)
}

type N2VocabularyRepositoryImpl struct {
	db *gorm.DB
}

func (i N2VocabularyRepositoryImpl) GetList(p *pkg.Paginator) ([]dao.N2VocabularySubject, error) {
	var iterms []dao.N2VocabularySubject

	var err = i.db.Scopes(p.GormPagination()).Find(&iterms).Error

	if err != nil {
		log.Error("Got an error finding n2 vocabulary", err)
		return nil, err
	}
	return iterms, nil
}

func (i N2VocabularyRepositoryImpl) GetQuestionById(p *pkg.Paginator, questionId int) ([]dao.N2VocabularySubjectOption, error) {
	var iterms []dao.N2VocabularySubjectOption

	var err = i.db.Where("questionId = ?", questionId).Find(&iterms).Error

	if err != nil {
		log.Error("Got an error finding n2 vocabulary", err)
		return nil, err
	}
	return iterms, nil
}

func N2VocabularyRepositoryInit(db *gorm.DB, paginator *pkg.Paginator) *N2VocabularyRepositoryImpl {
	return &N2VocabularyRepositoryImpl{
		db: db,
	}
}
