package yamapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type YamahaConfig struct {
	apiVer  string
	schema  string
	address string
	apiPath string
	zone    string
}

// TODO: move vars declaration to the API caller package
var (
	yReceiver = YamahaConfig{
		apiVer:  "v2",
		schema:  "http://",
		address: "192.168.1.17",
		apiPath: "/YamahaExtendedControl/",
		zone:    "main",
	}
	yClient = &http.Client{}
	baseUrl = yReceiver.schema + yReceiver.address + yReceiver.apiPath + yReceiver.apiVer
)

func SetVolume(volume byte) {

	req, err := http.NewRequest("GET", baseUrl+"/"+yReceiver.zone+"/setVolume?volume="+strconv.Itoa(int(volume)), nil)
	if err != nil {
		log.Fatalln(err)
	}

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

func SetPower(power string) {
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
