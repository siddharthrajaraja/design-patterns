package builder_pattern

type CarBuilder struct {
	seats       int
	wheels      int
	doors       int
	vehicleType string
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) setNumberOfWheels() {
	c.wheels = 4
}

func (c *CarBuilder) setNumberOfSeats() {
	c.seats = 5
}

func (c *CarBuilder) setNumberDoors() {
	c.doors = 4
}

func (c *CarBuilder) setVehicleType() {
	c.vehicleType = CarType
}

func (c *CarBuilder) getVehicle() *Vehicle {
	return NewVehicle(c.wheels, c.doors, c.seats, c.vehicleType)
}
