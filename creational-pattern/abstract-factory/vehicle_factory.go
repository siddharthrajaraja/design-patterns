package abstract_factory

import "errors"

type VehicleFactory interface {
	Build(vehicleType int) (Vehicle, error)
}

func BuildVehicleFactory(factoryType int) (VehicleFactory, error) {
	switch factoryType {
	case CarFactoryType:
		return NewCarFactory(), nil
	case BikeFactoryType:
		return NewBikeFactory(), nil
	default:
		return nil, errors.New("factory type in BuildVehicleFactory method mismatched")
	}
}
