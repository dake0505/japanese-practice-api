package vocabulary

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "n2_vocabulary_subject",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&N2VocabularySubject{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("n2_vocabulary_subject")
			},
		},
		{
			ID: "n2_vocabulary_subject_option",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&N2VocabularySubjectOption{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("n2_vocabulary_subject_option")
			},
		},
	}
}
