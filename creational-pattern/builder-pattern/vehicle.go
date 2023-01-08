package builder_pattern

type Vehicle struct {
	wheels      int
	doors       int
	seats       int
	vehicleType string
}

func NewVehicle(wheels, doors, seats int, vehicleType string) *Vehicle {
	return &Vehicle{
		wheels:      wheels,
		doors:       doors,
		seats:       seats,
		vehicleType: vehicleType,
	}
}
