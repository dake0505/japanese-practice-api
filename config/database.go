package config

import (
	"fmt"
	"gin-gonic-api/app/domain/dao/vocabulary"
	"gin-gonic-api/app/domain/dao/auth"
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
		allMigration = append(allMigration, vocabulary.Migration()...)
    allMigration = append(allMigration, auth.Migration()...)

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
