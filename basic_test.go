package main

// Basic imports
import (
	"github.com/kcmvp/idm/infra"
	"github.com/kcmvp/profile"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BasicTestSuite struct {
	suite.Suite
}

func (suite *BasicTestSuite) SetupSuite() {
	profile.With("test")
	infra.SetupClient()
}

func (suite *BasicTestSuite) TearDownSuite() {
}

func (suite *BasicTestSuite) TestBasic() {
	assert.Equal(suite.T(), 1, 1)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BasicTestSuite))
}
