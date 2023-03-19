package drivers

import (
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type Sensor struct {
	Type        string
	Unit        string
	Temperature float32
	Humidity    float32
}

var cachedTemperature, cachedHumidity float32
var cachedTime time.Time

func (sensor Sensor) GetValue(unit string, cacheTTL time.Duration) (Sensor, bool, error) {
	var err error
	sensor.Unit = unit
	if !time.Now().After(cachedTime.Add(cacheTTL)) {
		sensor.Temperature, sensor.Humidity = cachedTemperature, cachedHumidity
		return sensor, true, err
	}
	sht3x := i2c.NewSHT3xDriver(raspi.NewAdaptor())
	sht3x.Units = unit
	sht3x.Start()
	sensor.Temperature, sensor.Humidity, err = sht3x.Sample()
	if err == nil {
		cachedTemperature, cachedHumidity = sensor.Temperature, sensor.Humidity
		cachedTime = time.Now()
	}
	sht3x.Halt()
	return sensor, false, err
}
