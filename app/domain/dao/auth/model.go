package auth

import (
	"gin-gonic-api/app/domain/dao"
)


type Auth struct {
	ID       string `gorm:"column:id; primary_key; not null" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;->:false" json:"-"`
	dao.BaseModel
}

func (Auth) TableName() string {
	return "auth"
}
