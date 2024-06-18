package answeritem

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "answer_item",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&AnswerItem{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("answer_item")
			},
		},
	}
}
