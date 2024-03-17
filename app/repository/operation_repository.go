package repository

import (
	"gorm.io/gorm"
)

type OperationRepository interface {
	NextOperation() any
}

type OperationRepositoryImpl struct {
	db *gorm.DB
}

func (o OperationRepositoryImpl) NextOperation() any {
	return 1
}

func OperationRepositoryInit(db *gorm.DB) *OperationRepositoryImpl {
	return &OperationRepositoryImpl{
		db: db,
	}
}
