package dao

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
	CreatedBy gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
	UpdatedBy gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}
