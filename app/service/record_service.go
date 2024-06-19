package service

import "gin-gonic-api/app/repository"

type RecordService interface {
}

type RecordServiceImpl struct {
	recordRepository repository.RecordRepository
}

func RecordServiceInit(recordRepository repository.RecordRepository) *RecordServiceImpl {
	return &RecordServiceImpl{
		recordRepository: recordRepository,
	}
}
