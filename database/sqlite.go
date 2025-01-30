package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Sqlite(location string, logMode bool) (DB *gorm.DB, err error) {
	logLevel := logger.Warn
	if logMode {
		logLevel = logger.Info
	}

	return gorm.Open(sqlite.Open(location), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
}
