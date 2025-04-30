package config

import (
	"fmt"

	"github.com/g-celente/AuthService/src/app/database"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() (*gorm.DB, error) {
	var err error

	// Initialize Postgres
	db, err = database.InitializePostgres()

	if err != nil {
		return db, fmt.Errorf("error initializing sqlite: %v", err)
	}

	return db, nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
