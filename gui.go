package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

func initGui() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	// TODO: define widgets
	// help := NewHelpWidget("help", 1, 1, helpText)
	status := NewStatusbarWidget("volume", 1, 7, 50)
	// butdown := NewButtonWidget("butdown", 52, 7, "DOWN", statusDown(status))
	// butup := NewButtonWidget("butup", 58, 7, "UP", statusUp(status))
	g.SetManager(status)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}

func runGui() {
	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
