package core

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IAuthorRepository interface {
	Repository[Author]
}

type AuthorRepository struct {
	DB *gorm.DB
}

type Author struct {
	BaseTable
	Name        string        `gorm:"index" json:"name"`
	ExternalIds []ExternalIds `gorm:"serializer:json" json:"externalIds"`
}

func (repo *AuthorRepository) GetAll() ([]*Author, error) {
	var target []*Author
	repo.DB.Model(&Author{}).Scan(&target)
	if target == nil {
		return []*Author{}, nil
	}
	return target, nil
}

func (repo *AuthorRepository) GetByID(ids []string) ([]*Author, error) {
	var target []*Author
	repo.DB.Model(&Author{}).Where("id in ?", ids).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "author not found")
	}
	return target, nil
}

func (repo *AuthorRepository) GetByName(names []string) ([]*Author, error) {
	var target []*Author
	repo.DB.Model(&Author{}).Where("name in ?", names).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "author not found")
	}
	return target, nil
}

func (repo *AuthorRepository) Create(entity *Author) (*Author, error) {
	repo.DB.Create(&entity)
	return entity, nil
}
