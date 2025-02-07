package author

import (
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

func (suite *AuthorTestSuite) TestCreateAuthor() {
	t := suite.T()

	suite.Run("CreatesBookAndReturns", func() {
		out, _ := suite.repo.Create(&core.Author{
			Name: "Book One",
			ExternalIds: []core.ExternalIds{
				{System: "calibre", Value: "1"},
				{System: "koreader", Value: "2"},
			},
		})

		assert.NotNil(t, out)
		assert.NotNil(t, out.ID)
		assert.Equal(t, "Book One", out.Name)
		assert.Equal(t, []core.ExternalIds{
			{System: "calibre", Value: "1"},
			{System: "koreader", Value: "2"},
		}, out.ExternalIds)
	})
}
