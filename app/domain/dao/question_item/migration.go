package questionitem

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "question_item",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&QuestionItem{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("question_item")
			},
		},
	}
}
