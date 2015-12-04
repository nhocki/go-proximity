package proximity

import "github.com/mmcloughlin/geohash"

// IntervalFinder provides a way to override the way to look for geohashes
// intervals. Use `proximity.DefaultIntervalFinder` if you don't want to
// override it.
type IntervalFinder func(lat, lng, radius float64) []Int64arr

// DefaultIntervalFinder gets intervals from Redis given a radius. It will remove
// neighborhoods that are right next to each other to do less queries.
var DefaultIntervalFinder = intervalsFromRadius

// LocationSet is a set of named points inside a Redis database.
type LocationSet struct {
	Name           string
	client         Client
	IntervalFinder IntervalFinder
}

// Add adds a point <lat, lng> with name `name` to the location set.
//
// Usage:
//   set.Add("Place Name", lat, lng)
func (set *LocationSet) Add(name string, lat, lng float64) error {
	_, err := set.client.ZAdd(set.Name, encode(lat, lng), name)
	return err
}

// Near gets the points within <radius> from the <lat, lng> point.
func (set *LocationSet) Near(lat, lng, radius float64) ([]string, error) {
	if set.IntervalFinder == nil || &set.IntervalFinder == nil {
		set.IntervalFinder = DefaultIntervalFinder
	}
	intervals := set.IntervalFinder(lat, lng, radius)
	return query(set.client, set.Name, intervals)
}

// NewLocationSet returns a LocationSet
func NewLocationSet(name string, client Client) *LocationSet {
	return &LocationSet{
		Name:   name,
		client: client,
	}
}

func encode(lat, lng float64) float64 {
	return float64(geohash.EncodeIntWithPrecision(lat, lng, maxBitDepth))
}
