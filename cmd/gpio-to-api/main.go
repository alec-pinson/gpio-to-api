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
	sensor        drivers.Sensor
)

func main() {
	log.Println("Starting...")

	config = config.Load()

	switch gpioType := config.GPIOType; {
	case gpioType == "button":
		drivers.MonitorButton(config.GPIOPin, config.OnPress, config.URL)
	case gpioType == "gpio":

	case gpioType == "sht3x":

	}

	apiServer.Start()
}
