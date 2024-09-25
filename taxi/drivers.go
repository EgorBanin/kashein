package taxi

type Drivers []*Driver

func (d Drivers) Filter(t Trip) FilteredDrivers {
	filtered := make(FilteredDrivers, 0)
	for i := range d {
		if d[i].Rate(t) > 0 {
			filtered = append(filtered, d[i])
		}
	}

	return filtered
}

type FilteredDrivers []*Driver

func (d FilteredDrivers) Notify(t Trip) {

}
