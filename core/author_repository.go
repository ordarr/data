package core

import "gorm.io/gorm"

type AuthorRepository struct {
	DB *gorm.DB
}

type Authors []*Author

func (repo *AuthorRepository) GetAll() Authors {
	var target Authors
	repo.DB.Model(&Author{}).Scan(&target)
	return target
}

func (repo *AuthorRepository) GetById(id string) Author {
	var target Author
	repo.DB.Model(&Author{}).Where("id = ?", id).Scan(&target)
	return target
}

func (repo *AuthorRepository) GetByName(name string) Author {
	var target Author
	repo.DB.Model(&Author{}).Where("name = ?", name).Scan(&target)
	return target
}
