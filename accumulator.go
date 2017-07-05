package main

type accumulator struct {
	config accumulator_config
	conn   RedisConn
}

func NewAccumulator(c accumulator_config) *accumulator {
	conn, e := NewRedisConn("tcp", c.RedisAddress)
	_exit_if(e, c.RedisAddress)
	p := &accumulator{
		config: c,
		conn:   conn,
	}
	return p
}

func (p *accumulator) Stop() {
	p.conn.Close()
}

func (p *accumulator) Incr(key string) {
	p.conn.Hincrby(p.config.RedisHashSetName, key, 1)
}
