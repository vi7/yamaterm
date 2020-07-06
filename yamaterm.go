package main

import (
	"fmt"
	"net/http"
	"strconv"
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
var yReceiver = yamahaConfig{
	apiVer:  "v2",
	schema:  "http://",
	address: "192.168.1.17",
	apiPath: "/YamahaExtendedControl/",
	zone:    "main",
}

var yClient = &http.Client{}
var baseUrl = yReceiver.schema + yReceiver.address + yReceiver.apiPath + yReceiver.apiVer

func setVolume(volume byte) {

	req, err := http.NewRequest("GET", baseUrl+"/"+yReceiver.zone+"/setVolume?volume="+strconv.Itoa(int(volume)), nil)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: implement Yamaha events subscription via headers
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := yClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)

	defer resp.Body.Close()

	// Response body handling
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%s\n", body)

	// return
}

func setPower(power string) {
	req, err := http.NewRequest("GET", baseUrl+"/"+yReceiver.zone+"/setPower?power="+power, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := yClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)

	defer resp.Body.Close()
}

func main() {

	// TODO: implement basic tui for volume, power and song info

	// var volume byte = 20
	// fmt.Printf("Setting the volume to %d\n", volume)
	// setVolume(volume)

	fmt.Println("Toggling power")
	setPower("toggle")
}
