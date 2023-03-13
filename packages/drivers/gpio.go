package drivers

import (
	"strconv"

	"gobot.io/x/gobot/platforms/raspi"
)

func TurnOn(pin int) error {
	return raspi.NewAdaptor().DigitalWrite(strconv.Itoa(pin), byte(1))
}

func TurnOff(pin int) error {
	return raspi.NewAdaptor().DigitalWrite(strconv.Itoa(pin), byte(0))
}
