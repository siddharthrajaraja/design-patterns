package builder_pattern

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BuilderClientSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	director    *Director
	carBuilder  VehicleBuilder
	busBuilder  VehicleBuilder
	bikeBuilder VehicleBuilder
}

// suite initialisation
func TestSingletonSuite(t *testing.T) {
	suite.Run(t, new(BuilderClientSuite))
}

// runs before the entire test suite
func (suite *BuilderClientSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())

	// initialise director
	suite.director = NewDirector()

	// initialise car builder
	suite.carBuilder = getVehicleBuilder(CarType)

	// initialise bus builder
	suite.busBuilder = getVehicleBuilder(BusType)

	// initialise bike builder
	suite.bikeBuilder = getVehicleBuilder(BikeType)
}

// runs before each test
func (suite *BuilderClientSuite) SetupTest() {

}

func (suite *BuilderClientSuite) AfterTest() {
	suite.ctrl.Finish()
}

func (suite *BuilderClientSuite) TestCarBuilder() {
	// set car builder in director
	suite.director.setBuilder(suite.carBuilder)

	// create car vehicle via director
	actualCar := suite.director.createVehicle()

	expectedCar := &Vehicle{
		seats:       5,
		wheels:      4,
		doors:       4,
		vehicleType: CarType,
	}

	suite.Equal(expectedCar, actualCar)
}

func (suite *BuilderClientSuite) TestBusBuilder() {
	// set car builder in director
	suite.director.setBuilder(suite.busBuilder)

	// create car vehicle via director
	actualBus := suite.director.createVehicle()

	expectedBus := &Vehicle{
		seats:       50,
		wheels:      6,
		doors:       2,
		vehicleType: BusType,
	}

	suite.Equal(expectedBus, actualBus)
}

func (suite *BuilderClientSuite) TestBikeBuilder() {
	// set car builder in director
	suite.director.setBuilder(suite.bikeBuilder)

	// create car vehicle via director
	actualBike := suite.director.createVehicle()

	expectedBike := &Vehicle{
		seats:       2,
		wheels:      2,
		doors:       0,
		vehicleType: BikeType,
	}

	suite.Equal(expectedBike, actualBike)
}
