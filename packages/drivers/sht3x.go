package drivers

import (
	"log"
	"os"
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

var firstRun bool = true
var debugMode bool

var sht3x = i2c.NewSHT3xDriver(raspi.NewAdaptor())

var cachedTemperature, cachedHumidity float32
var cachedTime time.Time

func Debug(s string, v ...any) {
	if os.Getenv("DEBUG") == "true" {
		log.Printf(s, v...)
	}
}

func (sensor *Sensor) GetValue(unit string, cacheTTL time.Duration) (Sensor, bool, error) {
	var err error

	if firstRun {
		debugMode = os.Getenv("DEBUG") == "true"
		Debug("first run")
		firstRun = false
		sensor.Unit = unit
		sht3x.Units = unit
		err = sht3x.Start()
		if err != nil {
			log.Fatal(err)
		}
		Debug("first run complete")
	}

	if !time.Now().After(cachedTime.Add(cacheTTL)) {
		Debug("using cache (TTL %s)", cacheTTL)
		sensor.Temperature, sensor.Humidity = cachedTemperature, cachedHumidity
		Debug("cache used")
		return *sensor, true, err
	}

	// retry 3 times if an error occurs, sleep 1 second each time
	for i := 0; i < 3; i++ {
		Debug("getting temps")
		sensor.Temperature, sensor.Humidity, err = sht3x.Sample()
		if err == nil {
			cachedTemperature, cachedHumidity = sensor.Temperature, sensor.Humidity
			cachedTime = time.Now()
			Debug("setting cache")
			return *sensor, false, err
		}
		time.Sleep(1 * time.Second)
	}

	return *sensor, false, err
}
