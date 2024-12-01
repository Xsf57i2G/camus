package camus

import (
	"time"

	"github.com/Xsf57i2G/frame"
	"github.com/Xsf57i2G/monet"
)

type Engine struct {
	renderer *monet.Renderer
	refresh  time.Duration
	delta    time.Duration
}

func Run(root Node) error {
	if err := root.Init(); err != nil {
		return err
	}
	var ticker = time.NewTicker(time.Second / 60)
	defer ticker.Stop()
	for range ticker.C {
		if err := root.Update(time.Second / 60); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) Draw(w *frame.Frame) error {
	return nil
}
