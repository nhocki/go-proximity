package proximity

import (
	"fmt"

	redis "github.com/mediocregopher/radix.v2/pool"
	"github.com/ride/go-proximity/wrappers/radix"
)

func hErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Example() {
	client, err := redis.New("tcp", "127.0.0.1:6379", 10)
	hErr(err)

	wrapper := radix.Wrap(client)
	set := NewLocationSet(key, wrapper)

	err = set.Add("Toronto", 43.6667, -79.4167)
	hErr(err)

	err = set.Add("New York", 40.7143, -74.0060)
	hErr(err)

	results, err := set.Near(43.6687, -79.4167, 500)
	hErr(err)

	tmp, err := set.Near(43.6687, -79.4167, 500000)
	hErr(err)
	results = append(results, tmp...)
	fmt.Println(results)

	client.Cmd("DEL", key)

	// Output: [Toronto Toronto New York]
}
