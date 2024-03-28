package auth

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "auth",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&Auth{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("auth")
			},
		},
	}
}
