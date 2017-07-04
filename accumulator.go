package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type accumulator struct {
	config accumulator_config
}

func test() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
