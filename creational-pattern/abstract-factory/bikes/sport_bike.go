package bikes

type SportBike struct{}

func NewSportBike() *SportBike {
	return &SportBike{}
}

func (*SportBike) NumWheels() int {
	return 2
}

func (*SportBike) NumSeats() int {
	return 1
}

func (*SportBike) GetMotorbikeType() string {
	return SportsBikeLabel
}
