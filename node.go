package camus

import (
	"time"

	"github.com/Xsf57i2G/frame"
)

type Node interface {
	Init() error
	Update(dt time.Duration) error
	Draw(w *frame.Frame) error
}

type Root struct {
	Node
}

func (r *Root) Init() error {
	return nil
}

func (r *Root) Update(dt time.Duration) error {
	return nil
}

func (r *Root) Draw(w *frame.Frame) error {
	return nil
}
