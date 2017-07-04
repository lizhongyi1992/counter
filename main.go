package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("counter")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go r.Run() // listen and serve on 0.0.0.0:8080

	s := <-waiting_for_interrupt_chan()
	fmt.Println("quit when catch signal:", s)
}
