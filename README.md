# go-proximity

**This is a BETA. Interface might change a bit.**

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
	redis "github.com/mediocregopher/radix.v2/pool"
)

var redisConn *redis.Pool

func init() {
	redisConn, _ = redis.New("tcp", "127.0.0.1:6379", 10)
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

# Customizing

## Setting the Redis Client

Since this library doesn't want to force you to use one Redis library or another,
a `LocationSet` has a `Client` interface.

```go
type Client interface {
	ZAdd(set string, score float64, value string) (int64, error)
	ZRangeByScore(set string, from float64, to float64) ([]string, error)
}
```

To make this easier, an interface is provided for [Radix Pool][pool]:

```go
func WrapRadixClient(radixClient) *Client
```

## Getting Intervals

A `LocationSet` takes an `IntervalFinder` function that users can override to
match their needs. If no function is provided a default one will be used.

To override it, create a function that matches:

```go
type IntervalFinder func(lat, lng, radius float64) []Int64arr
```

`Int64arr` is a sortable `type Int64arr []int64`.

# Tests

**To run tests, you must have Redis running on the default port.** Tests will
use the `go-proximity:test-set` key to run. All tests cleanup after they're
done to prevent problems.

Run tests with:

```sh
go test -v
```

Copyright (c) 2015 ride group inc.

[pool]:https://godoc.org/github.com/mediocregopher/radix.v2/pool
