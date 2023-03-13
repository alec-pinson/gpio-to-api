package main

import (
	"log"

	"github.com/alec-pinson/gpio-to-api/packages/drivers"
)

var (
	config        Config
	apiServer     APIServer
	turnOnOffList []string = []string{"gpio"}
	getValueList  []string = []string{"sht3x"}
)

func main() {
	log.Println("Starting...")

	config = config.Load()

	if config.GPIOType == "button" {
		drivers.MonitorButton(config.GPIOPin, config.OnPress, config.URL)
	} else {
		apiServer.Start()
	}
}
