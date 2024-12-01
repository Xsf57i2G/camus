package camus

import (
	"time"

	"github.com/Xsf57i2G/frame"
	"github.com/Xsf57i2G/monet"
)

type Engine struct {
	renderer *monet.Renderer
	refresh  time.Duration
}

func New() *Engine {
	return &Engine{
		renderer: &monet.Renderer{},
		refresh:  16 * time.Millisecond,
	}
}

func (e *Engine) Run(root Node) error {
	if err := root.Init(); err != nil {
		return err
	}
	return nil
}

func (e *Engine) Draw(w *frame.Frame) error {
	return nil
}
