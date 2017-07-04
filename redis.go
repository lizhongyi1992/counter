package main

import "github.com/garyburd/redigo/redis"

type RedisConnection interface {
	Close()
	Incr(key string)
	Hincrby(key, field string, increment int)
}

type redisconn struct {
	conn redis.Conn
}

func NewRedisConnection(network, address string) (RedisConnection, error) {
	p := &redisconn{}
	c, e := redis.Dial(network, address)
	if e != nil {
		return nil, e
	}
	p.conn = c
	return p, nil
}

func (p *redisconn) Close() {
	p.conn.Close()
}

func (p *redisconn) Incr(key string) {
	p.conn.Do("incr", key)
}

func (p *redisconn) Hincrby(key, field string, increment int) {
	p.conn.Do("hincrby", key, field, increment)
}
