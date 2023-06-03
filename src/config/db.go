package config

import (
	"fmt"

	"github.com/gustavoteixeira8/url-shortener/src/utils"
	"gorm.io/driver/postgres"
)

func DatabaseConfig() postgres.Config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("POSTGRES_HOST"),
		utils.GetEnv("POSTGRES_USERNAME"),
		utils.GetEnv("POSTGRES_PASSWORD"),
		utils.GetEnv("POSTGRES_DATABASE"),
		utils.GetEnv("POSTGRES_PORT"),
	)

	return postgres.Config{
		DSN: dsn,
	}
}
