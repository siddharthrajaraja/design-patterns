package builder_pattern

type Director struct {
	builder VehicleBuilder
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) setBuilder(builder VehicleBuilder) {
	d.builder = builder
}

func (d *Director) createVehicle() *Vehicle {
	d.builder.setVehicleType()
	d.builder.setNumberDoors()
	d.builder.setNumberOfWheels()
	d.builder.setNumberOfSeats()
	return d.builder.getVehicle()
}
