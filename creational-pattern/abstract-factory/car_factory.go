package abstract_factory

import (
	"errors"
	"github.com/siddharthrajaraja/design-patterns/creational-pattern/abstract-factory/cars"
)

type CarFactory struct{}

func NewCarFactory() VehicleFactory {
	return &CarFactory{}
}

func (*CarFactory) Build(vehicleType int) (Vehicle, error) {
	switch vehicleType {
	case LuxuryCarType:
		return cars.NewLuxuryCar(), nil
	case FamilyCarType:
		return cars.NewFamilyCar(), nil
	}
	return nil, errors.New("car type mismatched")
}
