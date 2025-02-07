package author

import (
	"github.com/ordarr/data/core"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

func (suite *AuthorTestSuite) TestGetExternalIds() {
	t := suite.T()

	suite.Run("ReturnsExternalIds", func() {
		suite.populate()

		out, _ := suite.repo.GetAll()

		assert.NotNil(t, out)
		assert.Equal(t, []core.ExternalIds{{
			System: "calibre",
			Value:  "1",
		}, {
			System: "koreader",
			Value:  "2",
		}}, out[0].ExternalIds)
		assert.Equal(t, []core.ExternalIds{{
			System: "calibre",
			Value:  "2",
		}, {
			System: "koreader",
			Value:  "3",
		}}, out[1].ExternalIds)
	})
}
