package singleton_pattern

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SingletonSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	obj  Singleton
}

// suite initialisation
func TestSingletonSuite(t *testing.T) {
	suite.Run(t, new(SingletonSuite))
}

// runs before the entire test suite
func (suite *SingletonSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.obj = GetInstance()
}

// runs before each test
func (suite *SingletonSuite) SetupTest() {

}

func (suite *SingletonSuite) AfterTest() {
	suite.ctrl.Finish()
}

func (suite *SingletonSuite) TestObjectCreation() {
	suite.NotNil(suite.obj)
}

func (suite *SingletonSuite) TestIsObjectSimilar() {
	newObj := GetInstance()
	suite.Equal(suite.obj, newObj)
}

func (suite *SingletonSuite) TestAddOne() {
	suite.Equal(0, suite.obj.GetCount())
	suite.obj.AddOne()
	suite.Equal(1, suite.obj.GetCount())
}
