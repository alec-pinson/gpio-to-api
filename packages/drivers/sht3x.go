package drivers

import (
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type Sensor struct {
	Type        string
	Unit        string
	Temperature float32
	Humidity    float32
}

func (sensor Sensor) GetValue(unit string) (Sensor, error) {
	var err error
	// sensor.Type = config.GPIOType
	sht3x := i2c.NewSHT3xDriver(raspi.NewAdaptor())
	sht3x.Units = unit
	sensor.Unit = unit
	sht3x.Start()
	sensor.Temperature, sensor.Humidity, err = sht3x.Sample()
	sht3x.Halt()
	return sensor, err
}
