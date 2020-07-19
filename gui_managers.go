package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
)

type StatusbarWidget struct {
	name string
	x, y int
	w    int
	val  float64
}

func NewStatusbarWidget(name string, x, y, w int) *StatusbarWidget {
	return &StatusbarWidget{name: name, x: x, y: y, w: w}
}

func (w *StatusbarWidget) SetVal(val float64) error {
	if val < 0 || val > 1 {
		return errors.New("invalid value")
	}
	w.val = val
	return nil
}

func (w *StatusbarWidget) Val() float64 {
	return w.val
}

func (w *StatusbarWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2, 0)
	if err != nil && !gocui.IsUnknownView(err) {
		return err
	}
	v.Clear()

	rep := int(w.val * float64(w.w-1))
	fmt.Fprint(v, strings.Repeat("▒", rep))
	return nil
}
