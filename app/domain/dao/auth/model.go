package auth

import (
	"gin-gonic-api/app/domain/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	ID       uint   `gorm:"column:id; primary_key; not null" json:"id"`
	UserId   string `gorm:"column:user_id;" json:"userId"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;" json:"password"`
	dao.BaseModel
}

func (Auth) TableName() string {
	return "auth"
}

func (item *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	item.UserId = uuid.New().String()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (item *Auth) BeforeUpdate(tx *gorm.DB) (err error) {
	item.UpdatedAt = time.Now()
	return
}
