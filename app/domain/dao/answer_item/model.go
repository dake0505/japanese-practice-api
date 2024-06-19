package answeritem

import (
	"gin-gonic-api/app/domain/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerItem struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	AnswerId   string `gorm:"column:answer_id; not null" json:"answerId"`
	AnswerDesc string `gorm:"column:answer_desc;" json:"answerDesc"`
	QuestionID string `gorm:"column:question_id;" json:"questionId"`
	dao.BaseModel
}

func (AnswerItem) TableName() string {
	return "answer_item"
}

func (item *AnswerItem) BeforeCreate(tx *gorm.DB) (err error) {
	item.AnswerId = uuid.New().String()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (item *AnswerItem) BeforeUpdate(tx *gorm.DB) (err error) {
	item.UpdatedAt = time.Now()
	return
}
