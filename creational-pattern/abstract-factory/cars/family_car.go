package cars

type FamilyCar struct {
}

func NewFamilyCar() *FamilyCar {
	return &FamilyCar{}
}

func (*FamilyCar) NumDoors() int {
	return 5
}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 5
}
