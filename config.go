package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

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
		Accumulators: []accumulator_config{
			accumulator_config{
				RedisAddress:        "localhost:6379",
				MysqlDB:             "localhost:3306",
				MysqlTable:          "test",
				FlushIntervalSecond: 60,
				MaxKeyCached:        10000,
			},
		},
	}
	return c
}

func gen_sample_config() string {
	out, e := yaml.Marshal(default_config())
	if e != nil {
		_err(e)
	}
	return string(out)
}

func ParseConfig(path string) (Config, error) {
	c := Config{}
	b, e := ioutil.ReadFile(path)
	if e != nil {
		return c, e
	}
	e = yaml.Unmarshal(b, &c)
	return c, e
}
