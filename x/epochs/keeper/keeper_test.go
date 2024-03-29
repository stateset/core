package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/stateset/core/app/apptesting"
	"github.com/stateset/core/x/epochs/types"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
	suite.Suite
	queryClient types.QueryClient
}

// Test helpers
func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
