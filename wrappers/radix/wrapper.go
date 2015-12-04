package radix

import (
	"github.com/fzzy/radix/redis"
	"github.com/ride/go-proximity/encoders/int_hash"
)

// EncodeFunc is a function that takes a lat, lng and converts it into a geohash.
// If no encoder is specified, `intHash.Encode` will be used by default.
type EncodeFunc func(lat, lng float64) float64

// Client that implements the proximity.Client interface
type Client struct {
	client  *redis.Client
	Encoder EncodeFunc
}

// Wrap takes a radix client and converts it a proximity.Client
func Wrap(c *redis.Client) *Client {
	return &Client{client: c}
}

// Add stores a point with a given name.
func (w *Client) Add(set, name string, lat, lng float64) error {
	if w.Encoder == nil {
		w.Encoder = intHash.Encode
	}

	_, err := w.client.Cmd("zadd", set, w.Encoder(lat, lng), name).Int64()
	return err
}

// Near queries redis and returns an array of stored locations.
// TODO: Make this a piped call.
func (w *Client) Near(set string, lat, lng, radius float64) ([]string, error) {
	intervals := intHash.IntervalsFromRadius(lat, lng, radius)

	var results []string
	for _, interval := range intervals {
		from, to := float64(interval[0]), float64(interval[1])

		responses, err := w.client.Cmd("ZRANGEBYSCORE", set, from, to).List()
		if err != nil {
			return nil, err
		}
		results = append(results, responses...)
	}
	return results, nil
}
