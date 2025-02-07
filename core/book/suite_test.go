package book

import (
	"context"
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type BookTestSuite struct {
	suite.Suite
	ctx      context.Context
	repo     core.BookRepository
	populate func() []*core.Book
}

func insertTestBooks(db *gorm.DB, ctx context.Context) []*core.Book {
	bookOne := &core.Book{
		ExternalIds: []core.ExternalIds{{
			System: "calibre",
			Value:  "1",
		}, {
			System: "koreader",
			Value:  "2",
		},
		},
		Title: "Book One",
	}
	bookTwo := &core.Book{
		ExternalIds: []core.ExternalIds{{
			System: "calibre",
			Value:  "2",
		}, {
			System: "koreader",
			Value:  "3",
		},
		},
		Title: "Book Two",
	}
	session := db.Session(&gorm.Session{Context: ctx})
	session.Create(&bookOne)
	session.Create(&bookTwo)

	return []*core.Book{
		bookOne, bookTwo,
	}
}

func (suite *BookTestSuite) SetupSubTest() {
	suite.ctx = context.Background()

	_db := core.Connect(&core.Config{
		Type:    "sqlite",
		Name:    "ordarr.db",
		LogMode: true,
	})

	suite.populate = func() []*core.Book {
		return insertTestBooks(_db, suite.ctx)
	}

	suite.repo = core.BookRepository{DB: _db}
}

func (suite *BookTestSuite) TearDownSubTest() {
	_ = os.Remove("ordarr.db")
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
