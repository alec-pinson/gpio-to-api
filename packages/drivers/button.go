package drivers

import (
	"log"
	"net/http"
	"strconv"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func MonitorButton(pin int, onPress string, url string) {
	button := gpio.NewButtonDriver(raspi.NewAdaptor(), strconv.Itoa(pin))
	button.On(gpio.ButtonRelease, func(data interface{}) {
		buttonPress(onPress, url)
	})
	button.Start()
	log.Printf("Monitoring button")
	// wait forever
	select {}
}

func buttonPress(onPress string, url string) {
	log.Printf("Button pressed, %v - %v", onPress, url)
	if onPress == "callUrl" {
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		} else {
			log.Print(resp)
		}
	}
}
