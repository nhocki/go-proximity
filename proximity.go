package proximity

// Client is an for a Redis client.
type Client interface {
	ZAdd(set string, score float64, value string) (int64, error)
	ZRangeByScore(set string, from float64, to float64) ([]string, error)
}
