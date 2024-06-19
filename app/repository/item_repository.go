package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_item"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemList() []dao.QuestionItem
	CreateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
	UpdateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
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

func (i ItemRepositoryImpl) CreateQuestionItem(item *dao.QuestionItem) dao.QuestionItem {
	var err = i.db.Create(item).Error
	if err != nil {

	}
	return *item
}

func (i ItemRepositoryImpl) UpdateQuestionItem(item *dao.QuestionItem) dao.QuestionItem {
	var err = i.db.Save(item).Error
	if err != nil {

	}
	return *item
}

func ItemRepositoryInit(db *gorm.DB) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{
		db: db,
	}
}
