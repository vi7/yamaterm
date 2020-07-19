package main

import "github.com/awesome-gocui/gocui"

var (
	g *gocui.Gui
)

func main() {

	// TODO: implement basic tui for volume, power and song info

	// var volume byte = 21
	// log.Printf("Setting the volume to %d\n", volume)
	// yamapi.SetVolume(volume)

	// log.Println("Toggling power")
	// setPower("toggle")

	initGui()
	runGui()

}
