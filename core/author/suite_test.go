package author

import (
	"context"
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type AuthorTestSuite struct {
	suite.Suite
	ctx      context.Context
	repo     core.AuthorRepository
	populate func() []*core.Author
}

func insertTestAuthors(db *gorm.DB, ctx context.Context) []*core.Author {
	authorOne := &core.Author{

		ExternalIds: []core.ExternalIds{{
			System: "calibre",
			Value:  "1",
		}, {
			System: "koreader",
			Value:  "2",
		},
		},
		Name: "Name One",
	}
	authorTwo := &core.Author{
		ExternalIds: []core.ExternalIds{{
			System: "calibre",
			Value:  "2",
		}, {
			System: "koreader",
			Value:  "3",
		},
		},
		Name: "Name Two",
	}
	session := db.Session(&gorm.Session{Context: ctx})
	session.Create(&authorOne)
	session.Create(&authorTwo)

	return []*core.Author{
		authorOne, authorTwo,
	}
}

func (suite *AuthorTestSuite) SetupSubTest() {
	suite.ctx = context.Background()

	_db := core.Connect(&core.Config{
		Type:    "sqlite",
		Name:    "ordarr.db",
		LogMode: true,
	})

	suite.populate = func() []*core.Author {
		return insertTestAuthors(_db, suite.ctx)
	}

	suite.repo = core.AuthorRepository{DB: _db}
}

func (suite *AuthorTestSuite) TearDownSubTest() {
	_ = os.Remove("ordarr.db")
}

func TestAuthorTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorTestSuite))
}
