package proximity

import "github.com/mmcloughlin/geohash"

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
	_, err := set.client.ZAdd(set.Name, encode(lat, lng), name)
	return err
}

func encode(lat, lng float64) float64 {
	return float64(geohash.EncodeIntWithPrecision(lat, lng, maxBitDepth))
}
