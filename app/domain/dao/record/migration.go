package record

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "record",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&Record{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("record")
			},
		},
	}
}
