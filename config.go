package main

type Config interface {
	Map() map[string]interface{}
}

type accumulator_config struct {
	RedisAddress  string
	RedisPassword string
	RedisDB       string

	MysqlAddress  string
	MysqlPassword string
	MysqlDB       string
	MysqlTable    string
	MysqlField    string

	FlushIntervalSecond int
	MaxKeyCached        int

	WriteDBTimeout int
	FailRetryTimes int
}
