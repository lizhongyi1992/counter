package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var genconf, show_version bool

	flag.BoolVar(&genconf, "genconf", false, "generate a sample config")
	flag.BoolVar(&show_version, "version", false, "show version string and exit")
	flag.Parse()

	if show_version {
		fmt.Println(version_string())
		os.Exit(0)
	}

	if genconf {
		fmt.Println(gen_sample_config())
		os.Exit(0)
	}

	config := ParseConfig()

	r := gin.Default()

	r.GET("/ping", ping_test)
	r.POST("/incr_view_count", incr_view_count)

	go r.Run(config.ListenAt)

	s := <-waiting_for_interrupt_chan()
	_log("quit when catch signal:", s)
}
