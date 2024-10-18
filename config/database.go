package config

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"os"
	"strconv"
	"time"
)

const (
	DB_NAME string = "go_rest_template"
)

type DatabaseInformation struct {
	DataSourceName  string
	MaxConnOpenned  int
	MaxConnIddle    int
	MaxConnLifetime time.Duration
}

func ConnectDB() *sqlx.DB {
	var err error

	dbInfo := GetDatabaseInformation()

	db, err := sqlx.Open("postgres", dbInfo.DataSourceName)
	if err != nil {
		logrus.Error("error while connecting to database, Error: ", err)
	}

	db.SetMaxOpenConns(dbInfo.MaxConnOpenned)
	db.SetMaxIdleConns(dbInfo.MaxConnIddle)
	db.SetConnMaxLifetime(dbInfo.MaxConnLifetime)

	runMigrations(db)

	return db
}

func GetDatabaseInformation() *DatabaseInformation {
	dsn := os.Getenv("DB_DSN")
	maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		logrus.Fatal("ENV DB_MAX_OPEN_CONN should be an integer. Error: ", err)
	}
	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		logrus.Fatal("ENV DB_MAX_IDLE_CONN should be an integer. Error: ", err)
	}

	return &DatabaseInformation{
		DataSourceName:  dsn,
		MaxConnOpenned:  maxOpen,
		MaxConnIddle:    maxIdle,
		MaxConnLifetime: time.Hour,
	}
}

func runMigrations(db *sqlx.DB) {
	workingDir, err := os.Getwd()
	if err != nil {
		panic("could not get current directory, " + err.Error())
	}

	postgresDriver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		panic("failed to load postgres driver, " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/config/migrations", workingDir),
		DB_NAME,
		postgresDriver,
	)
	if err != nil {
		panic("failed to setup migrations config, " + err.Error())
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			return
		}
		panic("failed to run migrations, " + err.Error())
	}

	logrus.Info("migrations ran sucessfuly")
}
