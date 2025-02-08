package user

import (
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

func (suite *UserTestSuite) TestCreateUser() {
	t := suite.T()

	suite.Run("CreatesUserAndReturns", func() {
		out, _ := suite.repo.Create(&core.User{
			Username: "User One",
		})

		assert.NotNil(t, out)
		assert.NotNil(t, out.ID)
		assert.Equal(t, "User One", out.Username)
	})
}
