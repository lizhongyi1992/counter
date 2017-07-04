package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

func _log(v ...interface{}) {
	log.Println("INF", v)
}

func waiting_for_interrupt_chan() chan os.Signal {

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	return c
}
