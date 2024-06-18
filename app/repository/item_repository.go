package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_item"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemList() []dao.QuestionItem
}

type ItemRepositoryImpl struct {
	db *gorm.DB
}

func (i ItemRepositoryImpl) GetItemList() []dao.QuestionItem {
	var items []dao.QuestionItem
	var err = i.db.Find(&items).Error
	if err != nil {
	}
	return items
}

func ItemRepositoryInit(db *gorm.DB) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{
		db: db,
	}
}
