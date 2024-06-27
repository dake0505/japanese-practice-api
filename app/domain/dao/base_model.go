package dao

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	// DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
	CreatedBy string `gorm:"column:created_by" json:"createdBy"`
	UpdatedBy string `gorm:"column:deleted_at" json:"updatedBy"`
}
