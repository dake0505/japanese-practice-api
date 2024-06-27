package repository

import (
	"gin-gonic-api/app/domain/dao/record"

	"gorm.io/gorm"
)

type RecordRepository interface {
	CreateRecord(item *record.Record, createdBy string) (record.Record, error)
	QueryRecord(item *record.Record, createdBy string) (record.Record, error)
	QueryRecordList(item *record.Record, createdBy string) ([]record.Record, error)
	QueryRecordByQuestionId(questionId string, createdBy string) (record.Record, error)
	UpdateRecord(item *record.Record, createdBy string) (record.Record, error)
}

type RecordRepositoryImpl struct {
	db *gorm.DB
}

func (r RecordRepositoryImpl) CreateRecord(item *record.Record, createdBy string) (record.Record, error) {
	item.CreatedBy = createdBy
	var err = r.db.Create(item).Error
	if err != nil {
	}
	return *item, nil
}

func (r RecordRepositoryImpl) QueryRecord(item *record.Record, createdBy string) (record.Record, error) {
	var result record.Record
	query := r.db.Where("created_by = ?", createdBy)
	if item.ID != 0 {
		query = query.Where("id = ?", item.ID)
	}
	if item.QuestionId != "" {
		query = query.Where("question_id = ?", item.QuestionId)
	}
	err := query.First(&result).Error
	if err != nil {
	}
	return result, nil
}

func (r RecordRepositoryImpl) QueryRecordList(item *record.Record, createdBy string) ([]record.Record, error) {
	var results []record.Record
	query := r.db.Where("created_by = ?", createdBy)
	if item.ID != 0 {
		query = query.Where("id = ?", item.ID)
	}
	if item.QuestionId != "" {
		query = query.Where("question_id = ?", item.QuestionId)
	}
	if item.RecordType != "" {
		query = query.Where("record_type = ?", item.RecordType)
	}
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r RecordRepositoryImpl) QueryRecordByQuestionId(questionId string, createdBy string) (record.Record, error) {
	var result record.Record
	query := r.db.Where("created_by = ?", createdBy)
	query = query.Where("question_id = ?", questionId)
	err := query.First(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r RecordRepositoryImpl) UpdateRecord(item *record.Record, createdBy string) (record.Record, error) {
	var existingRecord record.Record

	err := r.db.Where("id = ?", item.ID).First(&existingRecord).Error
	if err != nil {
		return record.Record{}, err
	}

	existingRecord.RecordId = item.RecordId
	existingRecord.RecordType = item.RecordType
	existingRecord.QuestionId = item.QuestionId
	existingRecord.AnswerId = item.AnswerId
	existingRecord.UserAnswerId = item.UserAnswerId

	err = r.db.Save(&existingRecord).Error
	if err != nil {
		return record.Record{}, err
	}

	return existingRecord, nil
}

func RecordRepositoryInit(db *gorm.DB) *RecordRepositoryImpl {
	return &RecordRepositoryImpl{
		db: db,
	}
}
