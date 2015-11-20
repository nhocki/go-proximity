package proximity

import "github.com/stretchr/testify/mock"

type testClient struct {
	mock.Mock
	results  []string
	intReply int64
}

func (c *testClient) ZAdd(set string, score float64, value string) (int64, error) {
	args := c.Called(set, score, value)
	return int64(args.Int(0)), args.Error(1)
}

func (c *testClient) ZRangeByScore(set string, from, to float64) ([]string, error) {
	args := c.Called(set, from, to)
	return args.Get(0).([]string), args.Error(1)
}
