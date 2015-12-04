package radix

import (
	"testing"

	"github.com/fzzy/radix/redis"
	"github.com/stretchr/testify/assert"
)

const key = "go-proximity:test-set"

func dial(t *testing.T) {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	assert.Nil(t, err)
	wrapper = Wrap(c)
}

var (
	c       *redis.Client
	wrapper *Client
	err     error
)

func TestAdd(t *testing.T) {
	dial(t)
	wrapper.Add(key, "test-value", 1.0, -1.0)

	results, err := c.Cmd("ZRANGE", key, 0, -1).List()
	assert.Nil(t, err)
	assert.Equal(t, "test-value", results[0])
	c.Cmd("DEL", key, 0.0, -1.0)
}

func TestNear(t *testing.T) {
	dial(t)

	wrapper.Add(key, "test-value", 1.0, -1.0)
	results, err := wrapper.Near(key, 1.0, -1.0, 5000)
	assert.Nil(t, err)
	assert.Equal(t, "test-value", results[0])
	c.Cmd("DEL", key, 0.0, -1.0)
}
