package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func Test_redis(t *testing.T) {
	c, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()
}

func Test_redis2(t *testing.T) {
	c, e := NewRedisConnection("tcp", "localhost:6379")
	if e != nil {
		t.Error(e)
	}
	c.Incr("a")
}
