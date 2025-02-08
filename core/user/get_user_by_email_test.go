package user

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *UserTestSuite) TestGetUserByEmailReturnsUserWhenFound() {
	t := suite.T()

	suite.Run("ReturnsPopulatedUser", func() {
		suite.populate()

		out, _ := suite.repo.GetByEmail([]string{"user-one@example.com"})

		assert.NotNil(t, out)
		assert.Len(t, out, 1)
		assert.NoError(t, uuid.Validate(out[0].ID))
	})

	suite.Run("ErrorWhenUserDoesntExist", func() {
		_, err := suite.repo.GetByEmail([]string{"some-"})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, status.Error(codes.NotFound, "user not found"))
	})
}
