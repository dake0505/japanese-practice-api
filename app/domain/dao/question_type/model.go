package questiontype

import (
	"gin-gonic-api/app/domain/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionType struct {
	ID       uint   `gorm:"column:id; primary_key; not null" json:"id"`
	TypeId   string `gorm:"column:type_id;" json:"typeId"`
	TypeName string `gorm:"column:type_name;" json:"typeName"`
	dao.BaseModel
}

func (QuestionType) TableName() string {
	return "question_type"
}

func (item *QuestionType) BeforeCreate(tx *gorm.DB) (err error) {
	item.TypeId = uuid.New().String()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (item *QuestionType) BeforeUpdate(tx *gorm.DB) (err error) {
	item.UpdatedAt = time.Now()
	return
}
