package db

import (
	"fmt"
	"strings"

	"github.com/gustavoteixeira8/url-shortener/src/config"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _DB *gorm.DB

func SetupDatabase() {
	postgresConfig := config.DatabaseConfig()

	db, err := gorm.Open(postgres.New(postgresConfig))

	if err != nil {
		err = CreateDatabase()

		if err != nil {
			logrus.Fatalf("Error creating database: %v", err)
		}

		db, err = gorm.Open(postgres.New(postgresConfig))

		if err != nil {
			logrus.Fatalf("Error starting database: %v", err)
		}
	}

	_DB = db

	logrus.Info("Database connected")
}

func CreateDatabase() error {
	postgresConfig := config.DatabaseConfig()
	dsn := postgresConfig.DSN
	dsnSplit := strings.Split(dsn, " ")

	for i, val := range dsnSplit {
		if strings.Contains(val, "dbname=") {
			dsnSplit[i] = ""
		}
	}

	dsn = strings.Join(dsnSplit, " ")
	postgresConfig.DSN = dsn
	db, err := gorm.Open(postgres.New(postgresConfig))

	if err != nil {
		return err
	}

	query := `CREATE DATABASE %s;`
	queryFormatted := fmt.Sprintf(query, utils.GetEnv("POSTGRES_DATABASE"))

	err = db.Exec(queryFormatted).Error

	if err != nil {
		return err
	}

	logrus.Infof("Database %s created", utils.GetEnv("POSTGRES_DATABASE"))

	return nil
}

func GetDB() *gorm.DB {
	return _DB
}
