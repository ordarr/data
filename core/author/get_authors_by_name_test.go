package author

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *AuthorTestSuite) TestGetAuthorByName() {
	t := suite.T()

	suite.Run("ReturnsPopulatedAuthor", func() {
		suite.populate()

		out, _ := suite.repo.GetByName([]string{"Name One"})

		assert.NotNil(t, out)
		assert.Len(t, out, 1)
		assert.NoError(t, uuid.Validate(out[0].ID))
	})

	suite.Run("ErrorWhenAuthorDoesntExist", func() {
		_, err := suite.repo.GetByName([]string{"result shouldn't exist"})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, status.Error(codes.NotFound, "author not found"))
	})
}
