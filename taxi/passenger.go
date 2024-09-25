package taxi

type PassengerData struct {
	Id   string
	Name string
}

type Passenger struct {
	PassengerData
}

func NewPassenger(data PassengerData) *Passenger {
	return &Passenger{PassengerData: data}
}

func (p *Passenger) CreateTrip(data TripData) (*Trip, error) {
	data.Passenger = p.Id

	return NewTrip(data), nil
}
