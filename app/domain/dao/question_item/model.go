package questionitem

import "gin-gonic-api/app/domain/dao"

type QuestionItem struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	QuestionID    string `gorm:"column:id; not null" json:"questionId"`
	QuestionTitle string `gorm:"column: question_title; not null" json:"questionTitle"`
	AnswerId      string `gorm:"column: answer_id" json:"answerId"`
	dao.BaseModel
}

func (QuestionItem) TableName() string {
	return "question_item"
}
