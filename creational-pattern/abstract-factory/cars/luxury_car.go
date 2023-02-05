package cars

type LuxuryCar struct {
}

func NewLuxuryCar() *LuxuryCar {
	return &LuxuryCar{}
}

func (*LuxuryCar) NumDoors() int {
	return 4
}

func (*LuxuryCar) NumWheels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 5
}
