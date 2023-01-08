package builder_pattern

type BusBuilder struct {
	seats       int
	wheels      int
	doors       int
	vehicleType string
}

func NewBusBuilder() *BusBuilder {
	return &BusBuilder{}
}

func (c *BusBuilder) setNumberOfWheels() {
	c.wheels = 6
}

func (c *BusBuilder) setNumberOfSeats() {
	c.seats = 50
}

func (c *BusBuilder) setNumberDoors() {
	c.doors = 2
}

func (c *BusBuilder) setVehicleType() {
	c.vehicleType = BusType
}

func (c *BusBuilder) getVehicle() *Vehicle {
	return NewVehicle(c.wheels, c.doors, c.seats, c.vehicleType)
}
