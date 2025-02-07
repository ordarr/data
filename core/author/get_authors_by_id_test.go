package author

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *AuthorTestSuite) TestGetAuthorById() {
	t := suite.T()

	suite.Run("ReturnsPopulatedAuthor", func() {
		inserted := suite.populate()

		out, _ := suite.repo.GetById([]string{inserted[0].ID})

		assert.NotNil(t, out)
		assert.Len(t, out, 1)
		assert.Equal(t, inserted[0].Name, out[0].Name)
	})

	suite.Run("ErrorWhenAuthorDoesntExist", func() {
		t := suite.T()

		_, err := suite.repo.GetById([]string{"4783e133-d856-43f4-8d38-9e50c5996cad"})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, status.Error(codes.NotFound, "author not found"))
	})
}
