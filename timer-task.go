package main

import "time"

type TimerTask struct {
	stopchan  chan struct{}
	perSecond int
	f         func()
	running   bool
}

func NewTimerTask(perSecond int, f func()) *TimerTask {
	if perSecond < 1 {
		// limit
		perSecond = 1
	}
	p := &TimerTask{
		stopchan:  make(chan struct{}),
		perSecond: perSecond,
		f:         f,
		running:   false,
	}
	return p
}

func (p *TimerTask) Start() {
	p.running = true
	for p.running == true {
		_dbg(2)
		time.Sleep(1 * time.Second)
	}
}

func (p *TimerTask) Stop() {
	_dbg("aaa")
	close(p.stopchan)
	_dbg("bbb")
	p.running = false
}

func (p *TimerTask) StopChan() chan struct{} {
	return p.stopchan
}
