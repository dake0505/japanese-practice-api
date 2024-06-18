package repository

import (
	dao "gin-gonic-api/app/domain/dao/answer_item"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	CreateAnswerItem(item *dao.AnswerItem) dao.AnswerItem
	QueryAnswerList() []dao.AnswerItem
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

func AnswerRepositoryInit(db *gorm.DB) *AnswerRepositoryImpl {
	return &AnswerRepositoryImpl{
		db: db,
	}
}
