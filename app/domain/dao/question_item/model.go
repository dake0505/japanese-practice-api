package questionitem

import "gin-gonic-api/app/domain/dao"

type QuestionItem struct {
	ID string `gorm:"column:id; primary_key; not null" json:"id"`
	dao.BaseModel
}

func (QuestionItem) TableName() string {
	return "question_item"
}
