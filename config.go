package main

type Config struct {
	ListenAt     string
	Accumulators []accumulator_config
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

func default_config() Config {
	c := Config{
		ListenAt: ":8888",
	}
	return c
}

func gen_sample_config() string {
	return `ListenAt: :8888
`
}

func ParseConfig() Config {
	c := default_config()
	return c
}
