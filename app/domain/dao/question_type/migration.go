package questiontype

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "question_type",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&QuestionType{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("question_type")
			},
		},
	}
}
