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

func (sensor *Sensor) GetValue(unit string, cacheTTL time.Duration) (Sensor, bool, error) {
	var err error
	sensor.Unit = unit
	if !time.Now().After(cachedTime.Add(cacheTTL)) {
		sensor.Temperature, sensor.Humidity = cachedTemperature, cachedHumidity
		return *sensor, true, err
	}

	// retry 3 times if an error occurs, sleep 1 second each time
	for i := 0; i < 3; i++ {
		sht3x := i2c.NewSHT3xDriver(raspi.NewAdaptor())
		sht3x.Units = unit
		sht3x.Start()
		sensor.Temperature, sensor.Humidity, err = sht3x.Sample()
		sht3x.Halt()

		if err == nil {
			cachedTemperature, cachedHumidity = sensor.Temperature, sensor.Humidity
			cachedTime = time.Now()
			continue
		}
		time.Sleep(1 * time.Second)
	}

	return *sensor, false, err
}
