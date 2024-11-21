package core

import (
	"context"
	"fmt"
)

func ExampleNewRedis() {
	NewRedis("localhost:6379", "", 0)
	fmt.Printf("%T\n", G_REDIS)
	ping := G_REDIS.Ping(context.Background())
	fmt.Println(ping.Val())
	fmt.Println(ping.Err())
	// Output:
	// *redis.Client
	// PONG
	// <nil>
}
