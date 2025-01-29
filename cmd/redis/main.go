package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		ClientName: "go-experiment",
		Password:   "",
		DB:         0,
		Protocol:   3,
	})

	status := rdb.Set(ctx, "demo", "demo value", 0)
	if err := status.Err(); err != nil {
		fmt.Printf("Error when set operation: %v", err)
	}

	res := rdb.Get(ctx, "demo")
	val, err := res.Result()
	switch {
	case err == redis.Nil:
		fmt.Println("The key demo doesn't exist")
	case err != nil:
		fmt.Printf("Error when get result operation: %v", err)
	default:
		fmt.Printf("Result from key: %v\n", val)
	}

	res = rdb.GetSet(ctx, "demo-getset", "get-set-value")
	val, err = res.Result()
	// First execution of GetSet will return redis.Nil
	// next executions will return the value
	switch {
	case err == redis.Nil:
		fmt.Println("The key demo-getset doesn't exist")
	case err != nil:
		fmt.Printf("Error when get result operation: %v", err)
	default:
		fmt.Printf("Result from key: %v\n", val)
	}
}
