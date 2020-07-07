package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/awesome-gocui/gocui"
	// "github.com/awesome-gocui/gocui"
)

// TODO: implement adequate error handling
// TODO: extract API into the separate package

type yamahaConfig struct {
	apiVer  string
	schema  string
	address string
	apiPath string
	zone    string
}

// TODO: implement automatic discovery for the receiver
var (
	yReceiver = yamahaConfig{
		apiVer:  "v2",
		schema:  "http://",
		address: "192.168.1.17",
		apiPath: "/YamahaExtendedControl/",
		zone:    "main",
	}
	yClient = &http.Client{}
	baseUrl = yReceiver.schema + yReceiver.address + yReceiver.apiPath + yReceiver.apiVer
)

func setVolume(volume byte) {

	req, err := http.NewRequest("GET", baseUrl+"/"+yReceiver.zone+"/setVolume?volume="+strconv.Itoa(int(volume)), nil)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: implement Yamaha events subscription via headers
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := yClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Status)

	defer resp.Body.Close()

	// Response body handling
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s\n", body)

	// return
}

func setPower(power string) {
	req, err := http.NewRequest("GET", baseUrl+"/"+yReceiver.zone+"/setPower?power="+power, nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := yClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.Status)

	defer resp.Body.Close()
}

// TODO: gocui example, remove after tested

func guiManager(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("volume", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2, 0)
	if err != nil {

		if !gocui.IsUnknownView(err) {
			return err
		}

		fmt.Fprintln(v, "Hello world!")

		_, err := g.SetCurrentView("volume")
		if err != nil {
			return err
		}

	}
	return nil
}

func changeViewText(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	_, err := fmt.Fprintln(v, "   New text, yey!")
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {

	// TODO: implement basic tui for volume, power and song info

	// var volume byte = 21
	// log.Printf("Setting the volume to %d\n", volume)
	// setVolume(volume)

	// log.Println("Toggling power")
	// setPower("toggle")

	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(guiManager)

	err = g.SetKeybinding("", 'q', gocui.ModNone, quit)
	if err != nil {
		log.Panicln(err)
	}

	err = g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, changeViewText)
	if err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()
	if err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}

}
