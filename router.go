package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	video_acc *accumulator
}

func NewApp(config Config) *App {
	app := &App{}
	app.video_acc = NewAccumulator(config.VideoAcc)
	return app
}

func (p *App) incr_video_views(c *gin.Context) {
	var video_id int
	p.video_acc.Incr(fmt.Sprint(video_id))
	c.Status(200)
}
