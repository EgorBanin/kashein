package taxi

type DriverData struct {
	Name string
}

type Driver struct {
	DriverData
}

func NewDriver(data DriverData) *Driver {
	return &Driver{DriverData: data}
}

func (d *Driver) Rate(t Trip) float64 {
	return 1
}

func (d *Driver) Notify(t Trip) {

}
