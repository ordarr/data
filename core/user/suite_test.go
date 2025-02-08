package user

import (
	"context"
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	ctx      context.Context
	repo     core.UserRepository
	populate func() []*core.User
}

func insertTestUsers(db *gorm.DB, ctx context.Context) []*core.User {
	userOne := &core.User{
		Username: "User One",
		Email:    "user-one@example.com",
	}
	userTwo := &core.User{
		Username: "User Two",
		Email:    "user-two@example.com",
	}
	session := db.Session(&gorm.Session{Context: ctx})
	session.Create(&userOne)
	session.Create(&userTwo)

	return []*core.User{
		userOne, userTwo,
	}
}

func (suite *UserTestSuite) SetupSubTest() {
	suite.ctx = context.Background()

	_db := core.Connect(&core.Config{
		Type:    "sqlite",
		Name:    "ordarr.db",
		LogMode: true,
	})

	suite.populate = func() []*core.User {
		return insertTestUsers(_db, suite.ctx)
	}

	suite.repo = core.UserRepository{DB: _db}
}

func (suite *UserTestSuite) TearDownSubTest() {
	_ = os.Remove("ordarr.db")
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
