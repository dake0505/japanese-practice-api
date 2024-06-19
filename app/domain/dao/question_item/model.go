package questionitem

import (
	"gin-gonic-api/app/domain/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionItem struct {
	ID            uint   `gorm:"column:id; primary_key" json:"id"`
	QuestionID    string `gorm:"column:question_id; not null" json:"questionId"`
	QuestionTitle string `gorm:"column:question_title; not null" json:"questionTitle"`
	AnswerId      string `gorm:"column:answer_id" json:"answerId"`
	dao.BaseModel
}

func (QuestionItem) TableName() string {
	return "question_item"
}

func (item *QuestionItem) BeforeCreate(tx *gorm.DB) (err error) {
	item.QuestionID = uuid.New().String()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (item *QuestionItem) BeforeUpdate(tx *gorm.DB) (err error) {
	item.UpdatedAt = time.Now()
	return
}
