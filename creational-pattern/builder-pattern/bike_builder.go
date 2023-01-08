package builder_pattern

type BikeBuilder struct {
	seats       int
	wheels      int
	doors       int
	vehicleType string
}

func NewBikeBuilder() *BikeBuilder {
	return &BikeBuilder{}
}

func (c *BikeBuilder) setNumberOfWheels() {
	c.wheels = 2
}

func (c *BikeBuilder) setNumberOfSeats() {
	c.seats = 2
}

func (c *BikeBuilder) setNumberDoors() {
	c.doors = 0
}

func (c *BikeBuilder) setVehicleType() {
	c.vehicleType = BikeType
}

func (c *BikeBuilder) getVehicle() *Vehicle {
	return NewVehicle(c.wheels, c.doors, c.seats, c.vehicleType)
}
