package main

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

type accumulator struct {
	config accumulator_config
}

func test_redis() {
	c, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()
}
