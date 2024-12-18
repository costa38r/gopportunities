package config

import (
	"os"

	"github.com/costa38r/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	//CHECK IF DATABASES EXISTS

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file noit found, creating...")
		// create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//cREATE DATABASE AND CONNECT
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("sqlite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
