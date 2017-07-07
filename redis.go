package main

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConn interface {
	Close()
	Hincrby(key, field string, increment int)
	Rename(oldkey, newkey string)
	Del(key string)
	Hkeys(key string) []string
	Hget(key, hkey string) (string, error)
	Exists(key string) bool
}

type redisconn struct {
	pool *redis.Pool
}

func NewRedisConn(network, address string) (RedisConn, error) {
	_dbg(network, address)
	//TODO: poolzition
	p := &redisconn{}
	c, e := redis.Dial(network, address)
	if e != nil {
		return nil, e
	}
	// just for test
	c.Close()

	newpool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial(network, address) },
	}
	p.pool = newpool
	return p, nil
}

func (p *redisconn) Close() {
	e := p.pool.Close()
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Hincrby(key, field string, increment int) {
	conn := p.pool.Get()
	_, e := conn.Do("hincrby", key, field, increment)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Rename(oldkey, newkey string) {
	conn := p.pool.Get()
	_, e := conn.Do("rename", oldkey, newkey)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Del(key string) {
	conn := p.pool.Get()
	_, e := conn.Do("del", key)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Hkeys(key string) []string {
	conn := p.pool.Get()
	r := []string{}
	reply, e := conn.Do("hkeys", key)
	if e != nil {
		_err(e)
		return r
	}
	bytes, ok := reply.([]interface{})
	if !ok {
		return r
	}
	for _, v := range bytes {
		r = append(r, string(v.([]byte)))
	}
	return r
}

func (p *redisconn) Hget(key, hkey string) (string, error) {
	conn := p.pool.Get()
	reply, e := conn.Do("hget", key, hkey)
	if e != nil {
		_err(e)
		return "", e
	}
	_dbg(reply)
	switch t := reply.(type) {
	case []byte:
		return string(t), nil
	case nil:
		return "", errors.New("")
	}
	return "", errors.New("")
}

func (p *redisconn) Exists(key string) bool {
	conn := p.pool.Get()
	reply, e := conn.Do("exists", key)
	if e != nil {
		_err(e)
	}
	b := reply.(int64)
	if b == 0 {
		return false
	} else {
		return true
	}
}
