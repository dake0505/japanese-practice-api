package service

import (
	"gin-gonic-api/app/domain/dao/record"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/repository"
)

type RecordService interface {
	CreateRecord(item dto.CreateRecordDto) record.Record
	UpdateRecord(data dto.UpdateFavoriteDto, updateType string) (record.Record, error)
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
	data, err := r.recordRepository.CreateRecord(&input, item.CreatedBy)
	if err != nil {
	}
	return data
}

func (r RecordServiceImpl) UpdateRecord(data dto.UpdateFavoriteDto, updateType string) (record.Record, error) {
	var existingRecord record.Record
	if updateType == "favorite" {
		currentRecord, err := r.recordRepository.QueryRecordByQuestionId(data.QuestionId, data.CreatedBy)
		if err != nil {
			newRecord := record.Record{
				QuestionId: data.QuestionId,
				RecordType: "favorite",
			}
			createdRecord, err := r.recordRepository.CreateRecord(&newRecord, data.CreatedBy)
			if err != nil {
				return existingRecord, err
			}
			return createdRecord, nil
		}
		if currentRecord.RecordType == "favorite" {
			currentRecord.RecordType = "unFavorite"
		} else {
			currentRecord.RecordType = "favorite"
		}
		updateRecord, err := r.recordRepository.UpdateRecord(&currentRecord, data.CreatedBy)
		if err != nil {
		}
		return updateRecord, nil
	}
	return existingRecord, nil
}

func RecordServiceInit(recordRepository repository.RecordRepository) *RecordServiceImpl {
	return &RecordServiceImpl{
		recordRepository: recordRepository,
	}
}
