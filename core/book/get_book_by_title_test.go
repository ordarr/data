package book

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *BookTestSuite) TestGetBookByTitleReturnsBookWhenFound() {
	t := suite.T()

	suite.Run("ReturnsPopulatedBook", func() {
		suite.populate()

		out, _ := suite.repo.GetByTitle([]string{"Book One"})

		assert.NotNil(t, out)
		assert.Len(t, out, 1)
		assert.NoError(t, uuid.Validate(out[0].ID))
	})

	suite.Run("ErrorWhenBookDoesntExist", func() {
		_, err := suite.repo.GetByTitle([]string{"some-random-id"})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, status.Error(codes.NotFound, "book not found"))
	})
}
