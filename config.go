package main

type Config interface {
	Map() map[string]interface{}
}

type CounterConfig struct {
	data []byte
}

func (p *CounterConfig) Map() interface{} {
	return nil
}

type accumulator_config struct {
	RedisAddress string
	RedisSetname string

	MysqlAddress string
	MysqlDB      string
	MysqlTable   string
	MysqlField   string

	FlushIntervalSecond int
	MaxKeyCached        int

	WriteDBTimeout int
	FailRetryTimes int
}
