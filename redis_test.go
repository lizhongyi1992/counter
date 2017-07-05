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
	c, e := NewRedisConn("tcp", "localhost:6379")
	if e != nil {
		t.Error(e)
	}
	c.Incr("a")
}

func Test_redis_rename(t *testing.T) {
	c, e := NewRedisConn("tcp", "localhost:6379")
	if e != nil {
		t.Error(e)
	}
	c.Incr("abc")
	c.Rename("abc", "cba")
}

func Test_redis_hkeys(t *testing.T) {
	c, e := NewRedisConn("tcp", "localhost:6379")
	if e != nil {
		t.Error(e)
	}
	r := c.Hkeys("acc_views")
	t.Log(r)
}

func Test_redis_hget(t *testing.T) {
	c, e := NewRedisConn("tcp", "localhost:6379")
	if e != nil {
		t.Error(e)
	}
	r := c.Hget("acc_views", "abc")
	t.Log(r)
}
