package repository

import (
	"gorm.io/gorm"
)

type RecordRepository interface {
}

type RecordRepositoryImpl struct {
	db *gorm.DB
}

func RecordRepositoryInit(db *gorm.DB) *RecordRepositoryImpl {
	return &RecordRepositoryImpl{
		db: db,
	}
}
