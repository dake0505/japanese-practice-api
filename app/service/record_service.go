package service

import (
	"gin-gonic-api/app/domain/dao/record"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/repository"
)

type RecordService interface {
	CreateRecord(item dto.CreateRecordDto) record.Record
}

type RecordServiceImpl struct {
	recordRepository repository.RecordRepository
}

func (r RecordServiceImpl) CreateRecord(item dto.CreateRecordDto) record.Record {
	input := record.Record{
		QuestionId:   item.QuestionId,
		AnswerId:     item.AnswerId,
		UserAnswerId: item.UserAnswerId,
		RecordType:   item.RecordType,
	}
	data := r.recordRepository.CreateRecord(&input, item.CreatedBy)
	return data
}

func RecordServiceInit(recordRepository repository.RecordRepository) *RecordServiceImpl {
	return &RecordServiceImpl{
		recordRepository: recordRepository,
	}
}
