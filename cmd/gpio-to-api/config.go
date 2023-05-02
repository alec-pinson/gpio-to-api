package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Debug    bool
	GPIOType string
	GPIOPin  int
	Unit     string
	OnPress  string
	URL      string
	CacheTTL time.Duration
}

func (config Config) Load() Config {
	var err error

	log.Println("Loading configuration...")

	// debug
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		config.Debug = true
	} else {
		config.Debug = false
	}

	// gpio type
	config.GPIOType = os.Getenv("GPIO_TYPE")
	if config.GPIOType == "" {
		log.Fatalln("Environment variable 'GPIO_TYPE' must be set.")
	}
	if !(valueInList(config.GPIOType, turnOnOffList) || valueInList(config.GPIOType, getValueList) || config.GPIOType == "button" || config.GPIOType == "test") {
		log.Fatalf("Unknown GPIO_TYPE: %v", config.GPIOType)
	}

	// gpio pin
	if config.GPIOType == "gpio" || config.GPIOType == "button" {
		config.GPIOPin, err = strconv.Atoi(os.Getenv("GPIO_PIN"))
		if err != nil {
			log.Fatalln("Value for 'GPIO_PIN' is invalid.")
		}
	}

	// unit
	config.Unit = os.Getenv("UNIT")
	if config.Unit == "" {
		config.Unit = "C"
	}

	// cache time, sensor values are cached, when endpoint is called the value is not taken again if within cache time
	config.CacheTTL, _ = time.ParseDuration(os.Getenv("CACHE_TTL"))
	if config.CacheTTL == 0 {
		config.CacheTTL = 1 * time.Minute
	}

	if config.GPIOType == "button" {
		// on press
		config.OnPress = os.Getenv("ON_PRESS")

		if config.OnPress == "" {
			log.Fatalln("'ON_PRESS' must be set")
		}
		// url
		config.URL = os.Getenv("URL")
	}

	log.Println("Configuration loaded...")

	return config
}
