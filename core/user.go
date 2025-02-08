package core

import (
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	Repository[User]
	GetByEmail(email []string) (*User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

type User struct {
	BaseTable
	Username string `gorm:"index" json:"username"`
	Email    string `gorm:"index" json:"email"`
	Password []byte `json:"-"`
}

func (repo *UserRepository) GetAll() ([]*User, error) {
	var target []*User
	repo.DB.Preload(clause.Associations).Model(&User{}).Scan(&target)
	if target == nil {
		return []*User{}, nil
	}
	return target, nil
}

func (repo *UserRepository) GetByID(ids []string) ([]*User, error) {
	var target []*User
	repo.DB.Model(&User{}).Where("id in ?", ids).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return target, nil
}

func (repo *UserRepository) GetByName(names []string) ([]*User, error) {
	var target []*User
	repo.DB.Model(&User{}).Where("username in ?", names).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return target, nil
}

func (repo *UserRepository) GetByEmail(emails []string) ([]*User, error) {
	var target []*User
	repo.DB.Model(&User{}).Where("email in ?", emails).Scan(&target)
	if target == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return target, nil
}

func (repo *UserRepository) Create(entity *User) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword(entity.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	repo.DB.Create(&User{
		Username: entity.Username,
		Email:    entity.Email,
		Password: hashed,
	})
	return entity, nil
}
