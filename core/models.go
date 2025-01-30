package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ids struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Calibre  uint   `gorm:"index" json:"calibre"`
	Koreader uint   `gorm:"index" json:"koreader"`
}

func (b *Ids) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type Book struct {
	Ids
	Title string `gorm:"index" json:"title"`
}

type Author struct {
	Ids
	Name string `gorm:"index" json:"name"`
}
