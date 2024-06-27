package repository

import (
	dao "gin-gonic-api/app/domain/dao/question_item"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemList() []dao.QuestionItem
	CreateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
	UpdateQuestionItem(item *dao.QuestionItem) dao.QuestionItem
	QueryQuestionDetail(questionId string) dao.QuestionItem
	QueryPreviousQuestionId(currentQuestionId string) *string
	QueryNextQuestionId(currentQuestionId string) *string
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

func (i ItemRepositoryImpl) QueryQuestionDetail(questionId string) dao.QuestionItem {
	var questionItem dao.QuestionItem
	err := i.db.Table("question_item").
		Where("question_id = ?", questionId).
		First(&questionItem).Error
	if err != nil {
		return dao.QuestionItem{}
	}
	return questionItem
}

func (i ItemRepositoryImpl) QueryPreviousQuestionId(currentQuestionId string) *string {
	var currentId uint
	err := i.db.Table("question_item").
		Select("id").
		Where("question_id = ?", currentQuestionId).
		Scan(&currentId).Error
	if err != nil || currentId == 0 {
		return nil
	}

	var previousQuestionId string
	err = i.db.Table("question_item").
		Select("question_id").
		Where("id < ?", currentId).
		Order("id DESC").
		Limit(1).
		Scan(&previousQuestionId).Error
	if err != nil || previousQuestionId == "" {
		return nil
	}
	return &previousQuestionId
}

func (i ItemRepositoryImpl) QueryNextQuestionId(currentQuestionId string) *string {
	var currentId uint
	err := i.db.Table("question_item").
		Select("id").
		Where("question_id = ?", currentQuestionId).
		Scan(&currentId).Error
	if err != nil || currentId == 0 {
		return nil
	}
	var nextQuestionId string
	err = i.db.Table("question_item").
		Select("question_id").
		Where("id > ?", currentId).
		Order("id ASC").
		Limit(1).
		Scan(&nextQuestionId).Error
	if err != nil || nextQuestionId == "" {
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
