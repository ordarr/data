package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseTable struct {
	ID      string         `gorm:"primaryKey" json:"id"`
	Created int64          `json:"created"`
	Updated int64          `json:"updated"`
	Deleted gorm.DeletedAt `json:"deleted"`
}

func (b *BaseTable) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	now := time.Now().Unix()
	b.Created = now
	b.Updated = now
	return
}

func (b *BaseTable) BeforeUpdate(tx *gorm.DB) (err error) {
	b.Updated = time.Now().Unix()
	return
}

type ExternalIds struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type Repository[T interface{}] interface {
	GetAll() ([]*T, error)
	GetByID(ids []string) ([]*T, error)
	Create(entity *T) (*T, error)
}
