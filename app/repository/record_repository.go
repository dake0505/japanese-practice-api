package repository

import (
	"gin-gonic-api/app/domain/dao/record"

	"gorm.io/gorm"
)

type RecordRepository interface {
	CreateRecord(item *record.Record, createdBy string) record.Record
}

type RecordRepositoryImpl struct {
	db *gorm.DB
}

func (r RecordRepositoryImpl) CreateRecord(item *record.Record, createdBy string) record.Record {
	item.CreatedBy = createdBy
	var err = r.db.Create(item).Error
	if err != nil {
	}
	return *item
}

func RecordRepositoryInit(db *gorm.DB) *RecordRepositoryImpl {
	return &RecordRepositoryImpl{
		db: db,
	}
}
