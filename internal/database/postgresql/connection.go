package connection

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() (db *gorm.DB, err error) {
	var dsn string

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL != "" {
		dsn = databaseURL
	} else {
		dbUser := config.Get().Database.User
		dbPass := config.Get().Database.Pass
		dbHost := config.Get().Database.Host
		dbPort := config.Get().Database.Port
		dbName := config.Get().Database.Name

		if len(os.Args) > 1 && os.Args[1] == "migrate" {
			dbUser = config.Get().Database.UserMigration
			dbPass = config.Get().Database.PassMigration
		}

		if dbUser == "" {
			return nil, errors.New("Invalid environment variable 'DB_USER'. Configure it and try again")
		}
		if dbName == "" {
			return nil, errors.New("Invalid environment variable 'DB_NAME'. Configure it and try again")
		}

		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	var dbConfig gorm.Config
	if config.Get().Env == "production" || config.Get().Database.EnableLog {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		)
		dbConfig = gorm.Config{
			Logger:                 newLogger,
			SkipDefaultTransaction: true,
		}
	} else {
		dbConfig = gorm.Config{
			SkipDefaultTransaction: true,
		}
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &dbConfig)

	if err != nil {
		return nil, errors.New("[postgres.OpenConnection()] - Error on create postgres connection.\nPlease, try to see your environment variables or database connection...\nError: " + err.Error())
	}
	return db, nil
}
