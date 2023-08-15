package testify_suite

// Basic imports
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	One int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test

func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("SetupSuite()")
}

func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite()")
}

func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest()")
}

func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest()")
}

func (suite *ExampleTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest(%s, %s)\n", suiteName, testName)
}

func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest(%s, %s)\n", suiteName, testName)
}

func (suite *ExampleTestSuite) TestOne() {
	fmt.Println("TestOne()")
	assert.Equal(suite.T(), 1, 1)
}

func (suite *ExampleTestSuite) TestTwo() {
	fmt.Println("TestTwo()")
	assert.Equal(suite.T(), 2, 2)
}

func (suite *ExampleTestSuite) TestThree() {
	fmt.Println("TestThree()")
	assert.Equal(suite.T(), 3, 3)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
