package config

import (
	"fmt"
	answeritem "gin-gonic-api/app/domain/dao/answer_item"
	"gin-gonic-api/app/domain/dao/auth"
	questionitem "gin-gonic-api/app/domain/dao/question_item"
	questiontype "gin-gonic-api/app/domain/dao/question_type"
	"gin-gonic-api/app/domain/dao/record"
	"log"
	"os"
	"strconv"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB() *gorm.DB {
	var err error
	// dsn := os.Getenv("DB_DSN_PRD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	doMigrate, err := strconv.ParseBool(os.Getenv("DO_MIGRATE"))

	if err != nil {
		doMigrate = false
	}

	if doMigrate {
		allMigration := []*gormigrate.Migration{}
		// allMigration = append(allMigration, vocabulary.Migration()...)
		allMigration = append(allMigration, auth.Migration()...)
		allMigration = append(allMigration, questiontype.Migration()...)
		allMigration = append(allMigration, questionitem.Migration()...)
		allMigration = append(allMigration, answeritem.Migration()...)
		allMigration = append(allMigration, record.Migration()...)

		m := gormigrate.New(db, gormigrate.DefaultOptions, allMigration)
		if err := m.Migrate(); err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}
		log.Println("Migration did run successfully")
	} else {
		log.Println("Skipping migrations...")
	}

	return db
}
