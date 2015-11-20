# go-proximity

This library provides an easy way to do proximity queries to locations stored
in a Redis set.

It should be noted that the method used here is not the most precise,
but the query is pretty fast, and should be appropriate for most consumer
applications looking for this basic function.

```go
package main

import (
	"fmt"

	"github.com/fzzy/radix/redis"
	"github.com/ride/go-proximity"
	"github.com/ride/go-proximity/wrappers/radix"
)

var redisConn *redis.Client

func init() {
	redisConn, _ = redis.Dial("tcp", "127.0.0.1:6379")
}

func main() {
	set := &proximity.LocationSet{
		Name:   "go-prox-test",
		client: radix.Wrap(redisConn),
	}

	set.Add("Toronto", 43.6667, -79.4167)
	set.Add("Philadelphia", 39.9523, -75.1638)

	res, _ := set.Near(43.6687, -79.4167, 5000)
	fmt.Println(res)
}
```

## Setting the Redis Client

Since this library doesn't want to force you to use one Redis library or another,
a `LocationSet` has a `Client` interface.

```go
type Client interface {
	ZAdd(set string, score float64, value string) (int64, error)
	ZRangeByScore(set string, from float64, to float64) ([]string, error)
}
```

To make this easier, an interface is provided for Radix:

```go
func WrapRadixClient(radixClient) *Client
```
