package main

import "github.com/gin-gonic/gin"

type App struct {
	video_acc *accumulator
}

func NewApp(config Config) *App {
	app := &App{}
	app.video_acc = NewAccumulator(config.VideoAcc)
	return app
}

func (p *App) Stop() {
	p.video_acc.Stop()
}

func (p *App) incr_video_views(c *gin.Context) {
	video_id := c.Query("video_id")
	if video_id == "" {
		c.Status(200)
		return
	}
	p.video_acc.Incr(video_id)
	c.Status(200)
}
