package user

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (suite *UserTestSuite) TestGetUserById() {
	t := suite.T()

	suite.Run("ReturnsPopulatedUser", func() {
		inserted := suite.populate()

		out, _ := suite.repo.GetByID([]string{inserted[0].ID})

		assert.NotNil(t, out)
		assert.Len(t, out, 1)
		assert.Equal(t, inserted[0].Username, out[0].Username)
	})

	suite.Run("ErrorWhenUserDoesntExist", func() {
		t := suite.T()

		_, err := suite.repo.GetByID([]string{"4783e133-d856-43f4-8d38-9e50c5996cad"})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, status.Error(codes.NotFound, "user not found"))
	})
}
