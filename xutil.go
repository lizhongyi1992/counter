package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	VERSION   string
	BUILT     string
	GITHASH   string
	GOVERSION string
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
	log.SetFlags(log.LstdFlags)
}

func _dbg(v ...interface{}) {
	logger.Println("DBG", v)
}

func _err(v ...interface{}) {
	log.Println("ERR", v)
}

func _exit_if(err error, v ...interface{}) {
	if err != nil {
		log.Println("ERR Exit", err, v)
		os.Exit(-1)
	}
}

func _log(v ...interface{}) {
	log.Println("INF", v)
}

func _exit() {
	os.Exit(0)
}

func waiting_for_interrupt_chan() chan os.Signal {

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	return c
}

func version_string() string {
	return VERSION + " " + GITHASH + " " + BUILT + " " + GOVERSION
}

func ping_test(c *gin.Context) {
	c.String(200, "ok")
}
