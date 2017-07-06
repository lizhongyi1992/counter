package main

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

type RedisConn interface {
	Close()
	Incr(key string)
	Hincrby(key, field string, increment int)
	Rename(oldkey, newkey string)
	Del(key string)
	Hkeys(key string) []string
	Hget(key, hkey string) (string, error)
	Exists(key string) bool
}

type redisconn struct {
	conn redis.Conn
}

func NewRedisConn(network, address string) (RedisConn, error) {
	_dbg(network, address)
	p := &redisconn{}
	c, e := redis.Dial(network, address)
	if e != nil {
		return nil, e
	}
	p.conn = c
	return p, nil
}

func (p *redisconn) Close() {
	e := p.conn.Close()
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Incr(key string) {
	_, e := p.conn.Do("incr", key)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Hincrby(key, field string, increment int) {
	_, e := p.conn.Do("hincrby", key, field, increment)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Rename(oldkey, newkey string) {
	_, e := p.conn.Do("rename", oldkey, newkey)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Del(key string) {
	_, e := p.conn.Do("del", key)
	if e != nil {
		_err(e)
	}
}

func (p *redisconn) Hkeys(key string) []string {
	r := []string{}
	reply, e := p.conn.Do("hkeys", key)
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
	reply, e := p.conn.Do("hget", key, hkey)
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
	reply, e := p.conn.Do("exists", key)
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
