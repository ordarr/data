package core

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
)

func Open(config *Config) (*gorm.DB, error) {
	logLevel := logger.Warn
	if config.LogMode {
		logLevel = logger.Info
	}

	opts := &gorm.Config{Logger: logger.Default.LogMode(logLevel)}

	switch config.Type {
	case "sqlite":
		return gorm.Open(sqlite.Open(config.Name), opts)
	case "postgres":
		conn := fmt.Sprintf(
			"postgres://%s:%s@%s/%s?%s",
			config.User,
			config.Pass,
			net.JoinHostPort(config.Host, config.Port),
			config.Name,
			"sslmode=disable",
		)
		return gorm.Open(postgres.Open(conn), opts)
	}
	return nil, errors.New("invalid database type")
}

func Connect(config *Config) *gorm.DB {
	db, err := Open(config)
	if err != nil {
		log.Fatalf("Core database connection error: %v", err)
	}

	err = db.AutoMigrate(&Author{}, &Book{})
	if err != nil {
		log.Fatalf("Core migration connection error: %v", err)
	}

	return db
}
