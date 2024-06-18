package repository

import "gorm.io/gorm"

type AnswerRepository interface {
}

type AnswerRepositoryImpl struct {
	db *gorm.DB
}

func AnswerRepositoryInit(db *gorm.DB) *AnswerRepositoryImpl {
	return &AnswerRepositoryImpl{
		db: db,
	}
}
