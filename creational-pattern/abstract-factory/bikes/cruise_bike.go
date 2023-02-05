package bikes

type CruiseBike struct{}

func NewCruiseBike() *CruiseBike {
	return &CruiseBike{}
}

func (*CruiseBike) NumWheels() int {
	return 2
}

func (*CruiseBike) NumSeats() int {
	return 2
}

func (*CruiseBike) GetMotorbikeType() string {
	return CruiseBikeLabel
}
