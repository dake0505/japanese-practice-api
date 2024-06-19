package repository

import (
	dao "gin-gonic-api/app/domain/dao/answer_item"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	CreateAnswerItem(item *dao.AnswerItem) dao.AnswerItem
	QueryAnswerList() []dao.AnswerItem
	QueryAnswerListByQuestionId(questionId string) []dao.AnswerItem
}

type AnswerRepositoryImpl struct {
	db *gorm.DB
}

func (a AnswerRepositoryImpl) CreateAnswerItem(item *dao.AnswerItem) dao.AnswerItem {
	var err = a.db.Create(item).Error
	if err != nil {
	}
	return *item
}

func (a AnswerRepositoryImpl) QueryAnswerList() []dao.AnswerItem {
	var items []dao.AnswerItem
	var err = a.db.Find(&items).Error
	if err != nil {
	}
	return *&items
}

func (a AnswerRepositoryImpl) QueryAnswerListByQuestionId(questionId string) []dao.AnswerItem {
	var items []dao.AnswerItem
	err := a.db.Where("question_id = ?", questionId).Find(&items).Error
	if err != nil {
		return nil
	}
	return items
}

func AnswerRepositoryInit(db *gorm.DB) *AnswerRepositoryImpl {
	return &AnswerRepositoryImpl{
		db: db,
	}
}
