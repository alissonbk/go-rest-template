package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ConnectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")
	maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		logrus.Fatal("ENV DB_MAX_OPEN_CONN should be an integer. Error: ", err)
	}
	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		logrus.Fatal("ENV DB_MAX_IDLE_CONN should be an integer. Error: ", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: createDBLogger(),
	})
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

func createDBLogger() gormlogger.Interface {
	var logLevel gormlogger.LogLevel
	var ignoreNotFound bool
	var parameterizedQueries bool
	var colorful bool
	switch strings.ToUpper(os.Getenv("DB_LOG_LEVEL")) {
	case "PROD":
		logLevel = gormlogger.Warn
		ignoreNotFound = true
		parameterizedQueries = true
		colorful = false
	case "DEV":
		logLevel = gormlogger.Info
		ignoreNotFound = false
		parameterizedQueries = false
		colorful = true
	case "INFO":
		logLevel = gormlogger.Info
		ignoreNotFound = false
		parameterizedQueries = false
		colorful = true
	case "WARN":
		logLevel = gormlogger.Warn
		ignoreNotFound = false
		parameterizedQueries = false
	case "ERROR":
		logLevel = gormlogger.Error
		ignoreNotFound = true
		parameterizedQueries = false
	case "SILENT":
		logLevel = gormlogger.Silent
		ignoreNotFound = true
		parameterizedQueries = true
	default:
		logLevel = gormlogger.Error
		ignoreNotFound = true
		parameterizedQueries = true
	}

	return gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: ignoreNotFound,
			ParameterizedQueries:      parameterizedQueries,
			Colorful:                  colorful,
		},
	)
}
