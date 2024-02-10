package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func ConnectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")
	maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		log.Fatal("ENV DB_MAX_OPEN_CONN should be an integer. Error: ", err)
	}
	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		log.Fatal("ENV DB_MAX_IDLE_CONN should be an integer. Error: ", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database, Error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error while acquiring sql.DB from gorm lib, Error: ", err)
	}
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
