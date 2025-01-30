package calibre

import (
	"gorm.io/gorm"
	"log"
	"tarantini.io/ordarr/data/database"
)

var Calibre *gorm.DB

func Connect(location string, logMode bool) {
	db, err := database.Sqlite(location, logMode)
	if err != nil {
		log.Fatalf("Calibre database connection error: %v", err)
	}

	err = db.AutoMigrate(&Author{})
	if err != nil {
		log.Fatalf("Calibre migration error: %v", err)
	}

	Calibre = db
}
