package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_item"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemList() []dao.QuestionItem
	CreateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
	UpdateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
	QueryQuestionDetail(id uint) dao.QuestionItem
	QueryPreviousQuestionId(currentQuestionId uint) *uint
	QueryNextQuestionId(currentQuestionId uint) *uint
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

func (i ItemRepositoryImpl) QueryQuestionDetail(id uint) dao.QuestionItem {
	var questionItem dao.QuestionItem
	err := i.db.First(&questionItem, id).Error
	if err != nil {
		return dao.QuestionItem{}
	}
	return questionItem
}

func (i ItemRepositoryImpl) QueryPreviousQuestionId(currentQuestionId uint) *uint {
	var previousQuestionId uint
	err := i.db.Table("question_item").
		Select("id").
		Where("id < ?", currentQuestionId).
		Order("id DESC").
		Limit(1).
		Scan(&previousQuestionId).Error
	if err != nil || previousQuestionId == 0 {
		return nil
	}
	return &previousQuestionId
}

func (i ItemRepositoryImpl) QueryNextQuestionId(currentQuestionId uint) *uint {
	var nextQuestionId uint
	err := i.db.Table("question_item").
		Select("id").
		Where("id > ?", currentQuestionId).
		Order("id ASC").
		Limit(1).
		Scan(&nextQuestionId).Error
	if err != nil || nextQuestionId == 0 {
		return nil
	}
	return &nextQuestionId
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
