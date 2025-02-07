package core

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBookRepository interface {
	Repository[Book]
}

type BookRepository struct {
	DB *gorm.DB
}

type Book struct {
	BaseTable
	Name        string        `gorm:"index" json:"name"`
	ExternalIds []ExternalIds `gorm:"serializer:json" json:"externalIds"`
}

func (repo *BookRepository) GetAll() ([]*Book, error) {
	var target []*Book
	repo.DB.Preload(clause.Associations).Model(&Book{}).Scan(&target)
	if target == nil {
		return []*Book{}, nil
	}
	return target, nil
}

func (repo *BookRepository) GetByID(ids []string) ([]*Book, error) {
	var target []*Book
	repo.DB.Model(&Book{}).Where("id in ?", ids).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}
	return target, nil
}

func (repo *BookRepository) GetByName(names []string) ([]*Book, error) {
	var target []*Book
	repo.DB.Model(&Book{}).Where("name in ?", names).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}
	return target, nil
}

func (repo *BookRepository) Create(entity *Book) (*Book, error) {
	repo.DB.Create(entity)
	return entity, nil
}
