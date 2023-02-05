package abstract_factory

import (
	"github.com/golang/mock/gomock"
	"github.com/siddharthrajaraja/design-patterns/creational-pattern/abstract-factory/bikes"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VehicleFactorySuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	invalidFactory VehicleFactory
	carFactory     VehicleFactory
	bikeFactory    VehicleFactory
}

// suite initialisation
func TestVehicleFactorySuite(t *testing.T) {
	suite.Run(t, new(VehicleFactorySuite))
}

// runs before the entire test suite
func (suite *VehicleFactorySuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
}

// runs before each test
func (suite *VehicleFactorySuite) SetupTest() {

}

func (suite *VehicleFactorySuite) AfterTest() {
	suite.ctrl.Finish()
}

func (suite *VehicleFactorySuite) TestInvalidFactoryType() {
	var err error
	suite.invalidFactory, err = BuildVehicleFactory(InvalidFactoryType)
	suite.Nil(suite.invalidFactory)
	suite.NotNil(err)
	suite.Equal("factory type in BuildVehicleFactory method mismatched", err.Error())
}

func (suite *VehicleFactorySuite) TestCarFactoryTypeForLuxuryCar() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(CarFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(LuxuryCarType)
	suite.NotNil(vehicle)
	suite.Nil(err)
	suite.Equal(5, vehicle.NumSeats())
	suite.Equal(4, vehicle.NumWheels())
	car, ok := vehicle.(Car)
	suite.True(ok)
	suite.NotNil(car)
	suite.Equal(4, car.NumDoors())
}

func (suite *VehicleFactorySuite) TestCarFactoryTypeForFamilyCar() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(CarFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(FamilyCarType)
	suite.NotNil(vehicle)
	suite.Nil(err)
	suite.Equal(5, vehicle.NumSeats())
	suite.Equal(4, vehicle.NumWheels())
	car, ok := vehicle.(Car)
	suite.True(ok)
	suite.NotNil(car)
	suite.Equal(5, car.NumDoors())
}

func (suite *VehicleFactorySuite) TestCarFactoryTypeForInvalidCarType() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(CarFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(InvalidCarType)
	suite.Nil(vehicle)
	suite.NotNil(err)
	suite.Equal("car type mismatched", err.Error())
}

func (suite *VehicleFactorySuite) TestBikeFactoryTypeForCruiseBike() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(BikeFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(CruiseBikeType)
	suite.NotNil(vehicle)
	suite.Nil(err)
	suite.Equal(2, vehicle.NumSeats())
	suite.Equal(2, vehicle.NumWheels())
	car, ok := vehicle.(Motorbike)
	suite.True(ok)
	suite.NotNil(car)
	suite.Equal(bikes.CruiseBikeLabel, car.GetMotorbikeType())
}

func (suite *VehicleFactorySuite) TestBikeFactoryTypeForSportsBike() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(BikeFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(SportBikeType)
	suite.NotNil(vehicle)
	suite.Nil(err)
	suite.Equal(1, vehicle.NumSeats())
	suite.Equal(2, vehicle.NumWheels())
	car, ok := vehicle.(Motorbike)
	suite.True(ok)
	suite.NotNil(car)
	suite.Equal(bikes.SportsBikeLabel, car.GetMotorbikeType())
}

func (suite *VehicleFactorySuite) TestBikeFactoryTypeForInvalidBikeType() {
	var err error
	suite.carFactory, err = BuildVehicleFactory(BikeFactoryType)
	suite.Nil(err)
	suite.NotNil(suite.carFactory)
	vehicle, err := suite.carFactory.Build(InvalidBikeType)
	suite.Nil(vehicle)
	suite.NotNil(err)
	suite.Equal("bike type mismatched", err.Error())
}
