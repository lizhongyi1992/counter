package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	var genconf, show_version bool
	var config_file string

	flag.BoolVar(&genconf, "genconf", false, "generate a sample config")
	flag.BoolVar(&show_version, "version", false, "show version string and exit")
	flag.StringVar(&config_file, "c", "config.yaml", "config file path")
	flag.Parse()

	if show_version {
		fmt.Println(version_string())
		_exit()
	}

	if genconf {
		fmt.Println(gen_sample_config())
		_exit()
	}

	config, err := ParseConfig(config_file)
	_exit_if(err)
	_log(config)

	r := gin.Default()

	app := NewApp(config)

	r.GET("/ping", ping_test)
	r.POST("/incr_video_views", app.incr_video_views)

	syncer := NewTimerTask(config.VideoAcc.FlushIntervalSecond, func() {})
	go syncer.Start()

	go r.Run(config.ListenAt)

	s := <-waiting_for_interrupt_chan()
	_log("quit when catch signal:", s)
	syncer.Stop()
	_log("waiting for syncer exit")
	<-syncer.StopChan()
}
