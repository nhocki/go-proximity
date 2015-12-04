package radix

import (
	redis "github.com/mediocregopher/radix.v2/pool"
)

// Client that implements the proximity.Client interface
type Client struct {
	client *redis.Pool
}

// Wrap takes a radix pool and converts it a proximity.Client
func Wrap(c *redis.Pool) *Client {
	return &Client{client: c}
}

// ZAdd adds a value to a sorted set with a given score.
func (w *Client) ZAdd(set string, score float64, value string) (int64, error) {
	c := w.client
	r := c.Cmd("ZADD", set, score, value)
	return r.Int64()
}

// ZRangeByScore gets all values from a sorted set in the [from, to] score.
func (w *Client) ZRangeByScore(set string, from float64, to float64) ([]string, error) {
	c := w.client
	r := c.Cmd("ZRANGEBYSCORE", set, from, to)
	return r.List()
}
