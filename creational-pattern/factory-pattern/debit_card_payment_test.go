package factory_pattern

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DebitCardPMTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func TestDebitCardPMSuite(t *testing.T) {
	suite.Run(t, new(DebitCardPMTestSuite))
}

// runs before each test
func (suite *DebitCardPMTestSuite) SetupTest() {

}

func (suite *DebitCardPMTestSuite) AfterTest() {
	suite.ctrl.Finish()
}

func (suite *DebitCardPMTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *DebitCardPMTestSuite) TestDebitCardPMWithRandomPaymentMethodType() {
	amount := 10.0
	paymentMethod, err := GetPaymentMehtod(RandomPaymentMethod, amount)
	suite.Nil(paymentMethod)
	suite.NotNil(err)
	suite.Equal("not implemented yet", err.Error())
}

func (suite *DebitCardPMTestSuite) TestCashPMWithCard() {
	amount := 15.0
	paymentMethod, err := GetPaymentMehtod(DebitCard, amount)
	suite.Nil(err)
	suite.NotNil(paymentMethod)
	suite.Equal("Paid 15.0 via debit_card", paymentMethod.Pay())
}
