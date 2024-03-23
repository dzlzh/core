package core

import (
	"context"
	"fmt"
)

func ExampleNewRedis() {
	r := NewRedis("localhost:6379", "", 0)
	fmt.Printf("%T\n", r)
	ping := r.Ping(context.Background())
	fmt.Println(ping.Val())
	fmt.Println(ping.Err())
	// Output:
	// *redis.Client
	// PONG
	// <nil>
}
