package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

var (
	g *gocui.Gui
)

func main() {
	// TODO: implement basic tui for volume, power and song info

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFrameColor = gocui.ColorRed

	help := NewHelpWidget("help", 1, 1, helpText)
	status := NewStatusbarWidget("status", 1, 7, 50)
	butdown := NewButtonWidget("butdown", 52, 7, "DOWN", statusDown(status))
	butup := NewButtonWidget("butup", 58, 7, "UP", statusUp(status))
	g.SetManager(help, status, butdown, butup)

	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Panicln(err)
	}

	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleButton)
	if err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()
	if err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}
