package factory_pattern

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CashPMTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func TestCashPMSuite(t *testing.T) {
	suite.Run(t, new(CashPMTestSuite))
}

// runs before each test
func (suite *CashPMTestSuite) SetupTest() {

}

func (suite *CashPMTestSuite) AfterTest() {
	suite.ctrl.Finish()
}

func (suite *CashPMTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *CashPMTestSuite) TestCashPMWithRandomPaymentMethodType() {
	amount := 10.0
	paymentMethod, err := GetPaymentMehtod(RandomPaymentMethod, amount)
	suite.Nil(paymentMethod)
	suite.NotNil(err)
	suite.Equal("not implemented yet", err.Error())
}

func (suite *CashPMTestSuite) TestCashPMWithCard() {
	amount := 15.0
	paymentMethod, err := GetPaymentMehtod(Cash, amount)
	suite.Nil(err)
	suite.NotNil(paymentMethod)
	suite.Equal("Paid 15.0 via cash", paymentMethod.Pay())
}
