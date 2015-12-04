package proximity

// LocationSet is a set of named points inside a Redis database.
type LocationSet struct {
	Name   string
	client Client
}

// Add adds a point <lat, lng> with name `name` to the location set.
//
// Usage:
//   set.Add("Place Name", lat, lng)
func (set *LocationSet) Add(name string, lat, lng float64) error {
	return set.client.Add(set.Name, name, lat, lng)
}

// Near gets the points within <radius> from the <lat, lng> point.
//
// Usage:
//
// set.Near(lat, lng, radius)
func (set *LocationSet) Near(lat, lng, radius float64) ([]string, error) {
	return set.client.Near(set.Name, lat, lng, radius)
}
