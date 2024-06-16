package questiontype

import "gin-gonic-api/app/domain/dao"

type QuestionType struct {
	ID   string `gorm:"column:id; primary_key; not null" json:"id"`
	Type string `gorm:"column:type;" json:"type"`
	dao.BaseModel
}

func (QuestionType) TableName() string {
	return "question_type"
}
