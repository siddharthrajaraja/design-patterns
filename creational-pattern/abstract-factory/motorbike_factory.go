package abstract_factory

import (
	"errors"
	"github.com/siddharthrajaraja/design-patterns/creational-pattern/abstract-factory/bikes"
)

type MotorbikeFactory struct{}

func NewBikeFactory() VehicleFactory {
	return &MotorbikeFactory{}
}

func (*MotorbikeFactory) Build(vehicleType int) (Vehicle, error) {
	switch vehicleType {
	case CruiseBikeType:
		return bikes.NewCruiseBike(), nil
	case SportBikeType:
		return bikes.NewSportBike(), nil
	}
	return nil, errors.New("bike type mismatched")
}
