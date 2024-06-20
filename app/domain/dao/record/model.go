package record

import (
	"gin-gonic-api/app/domain/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	ID           uint   `gorm:"column:id; primary_key; not null" json:"id"`
	RecordId     string `gorm:"column:record_id;" json:"recordId"`
	RecordType   string `gorm:"column:record_type;" json:"recordType"`
	QuestionId   string `gorm:"column:question_id" json:"questionId"`
	AnswerId     string `gorm:"column:answer_id" json:"answerId"`
	UserAnswerId string `gorm:"column:user_answer_id" json:"userAnswerId"`
	dao.BaseModel
}

func (Record) TableName() string {
	return "record"
}

func (item *Record) BeforeCreate(tx *gorm.DB) (err error) {
	item.RecordId = uuid.New().String()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (item *Record) BeforeUpdate(tx *gorm.DB) (err error) {
	item.UpdatedAt = time.Now()
	return
}
