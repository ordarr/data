package book

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

func (suite *BookTestSuite) TestGetAllBooks() {
	t := suite.T()

	suite.Run("ReturnsPopulatedList", func() {
		suite.populate()

		out, _ := suite.repo.GetAll()

		assert.NotNil(t, out)
		assert.Len(t, out, 2)
	})

	suite.Run("ReturnsEmptyList", func() {
		out, _ := suite.repo.GetAll()

		assert.NotNil(t, out)
		assert.Len(t, out, 0)
	})
}
