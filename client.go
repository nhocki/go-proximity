package proximity

// Client is an for a Redis client.
type Client interface {
	Add(set, name string, lat, lng float64) error
	Near(set string, lat, lng, radius float64) ([]string, error)
}
