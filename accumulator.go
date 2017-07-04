package main

type accumulator struct {
	config accumulator_config
}

func NewAccumulator(c accumulator_config) *accumulator {
	p := &accumulator{}
	return p
}

func (p *accumulator) Incr(key string) {
}
