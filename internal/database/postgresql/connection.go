package connection

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.luizalabs.com/taste-match-api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetConnection returns a configured instance of gorm connection
func GetConnection() (db *gorm.DB, err error) {
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
		err = errors.New("Invalid environment variable 'DB_USER'. Configure it and try again")
	}
	if dbPass == "" {
		err = errors.New("Invalid environment variable 'DB_PASS'. Configure it and try again")
	}
	if dbHost == "" {
		err = errors.New("Invalid environment variable 'DB_HOST'. Configure it and try again")
	}
	if dbPort == "" {
		err = errors.New("Invalid environment variable 'DB_PORT'. Configure it and try again")
	}
	if dbName == "" {
		err = errors.New("Invalid environment variable 'DB_NAME'. Configure it and try again")
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &dbConfig)

	if err != nil {
		return nil, errors.New("[connection.GetConnection()] - Error on create connection.\nPlease, try to see your environment variables or database connection...\nError: " + err.Error())
	}
	return db, nil
}
