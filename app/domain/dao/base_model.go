package dao

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
	// DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
	CreatedBy int `gorm:"->:false;column:created_by" json:"-"`
	UpdatedBy int `gorm:"->:false;column:deleted_at" json:"-"`
}
