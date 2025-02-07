package core

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBookRepository interface {
	Repository[Book]
	GetByTitle(name []string) (*Book, error)
}

type BookRepository struct {
	DB *gorm.DB
}

type Book struct {
	BaseTable
	Title       string        `gorm:"index" json:"title"`
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

func (repo *BookRepository) GetById(ids []string) ([]*Book, error) {
	var target []*Book
	repo.DB.Model(&Book{}).Where("id in ?", ids).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}
	return target, nil
}

func (repo *BookRepository) GetByTitle(titles []string) ([]*Book, error) {
	var target []*Book
	repo.DB.Model(&Book{}).Where("title in ?", titles).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}
	return target, nil
}

func (repo *BookRepository) Create(title string) (*Book, error) {
	book := &Book{Title: title}
	repo.DB.Create(book)
	return book, nil
}
