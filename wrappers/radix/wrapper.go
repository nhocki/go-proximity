package radix

import "github.com/fzzy/radix/redis"

// Client that implements the proximity.Client interface
type Client struct {
	client *redis.Client
}

// Wrap takes a radix client and converts it a proximity.Client
func Wrap(c *redis.Client) *Client {
	return &Client{client: c}
}

// ZAdd adds a value to a sorted set with a given score.
func (w *Client) ZAdd(set string, score float64, value string) (int64, error) {
	c := w.client
	r := c.Cmd("zadd", set, score, value)
	return r.Int64()
}

// ZRangeByScore gets all values from a sorted set in the [from, to] score.
func (w *Client) ZRangeByScore(set string, from float64, to float64) ([]string, error) {
	c := w.client
	r := c.Cmd("ZRANGEBYSCORE", set, from, to)
	return r.List()
}
