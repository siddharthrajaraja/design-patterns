package builder_pattern

type VehicleBuilder interface {
	setNumberOfWheels()
	setNumberDoors()
	setNumberOfSeats()
	setVehicleType()
	getVehicle() *Vehicle
}

func getVehicleBuilder(vehicleType string) VehicleBuilder {
	// return CarBuilder for Car type
	if vehicleType == CarType {
		return NewCarBuilder()
	}
	if vehicleType == BusType {
		return NewBusBuilder()
	}
	if vehicleType == BikeType {
		return NewBikeBuilder()
	}
	return nil
}
