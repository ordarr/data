package core

import "gorm.io/gorm"

type BookRepository struct {
	DB *gorm.DB
}

type Books []*Book

func (repo *BookRepository) GetAll() Books {
	var target Books
	repo.DB.Model(&Book{}).Scan(&target)
	return target
}

func (repo *BookRepository) GetById(id string) Book {
	var target Book
	repo.DB.Model(&Book{}).Where("id = ?", id).Scan(&target)
	return target
}

func (repo *BookRepository) GetByTitle(title string) Book {
	var target Book
	repo.DB.Model(&Book{}).Where("title = ?", title).Scan(&target)
	return target
}
