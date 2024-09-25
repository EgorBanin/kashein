package taxi

type TripData struct {
	Passenger     string
	Driver        string
	Price         int
	PaymentMethod string
}

type Trip struct {
	TripData
	ch chan TripData
}

func NewTrip(data TripData) *Trip {
	return &Trip{TripData: data}
}

func (t *Trip) ChangePrice(newPrice int) {

}
