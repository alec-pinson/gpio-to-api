package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/alec-pinson/gpio-to-api/packages/drivers"
)

type APIServer struct{}

func (apiServer APIServer) Start() {
	var wg sync.WaitGroup
	log.Println("Starting API server...")
	http.HandleFunc("/", apiServer.Endpoint)
	wg.Add(1)
	go http.ListenAndServe(":8080", nil)
	log.Println("API Server started...")
	wg.Wait()
}

func (apiServer APIServer) Endpoint(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; {
	case path == "":
		apiServer.getValue(w)
	case path == "turnOn":
		apiServer.turnOn(w)
	case path == "turnOff":
		apiServer.turnOff(w)
	case path == "health/live" || path == "health/ready":
		fmt.Fprintf(w, "ok")
	}
}

func (apiServer APIServer) turnOn(w http.ResponseWriter) {
	if !valueInList(config.GPIOType, turnOnOffList) {
		writeErrorResponse(w, "GPIO type cannot be turned on")
		return
	}
	err := drivers.TurnOn(config.GPIOPin)
	if err != nil {
		writeErrorResponse(w, "%s", err)
		return
	} else {
		writeSuccessResponse(w, "Turned on")
	}
}

func (apiServer APIServer) turnOff(w http.ResponseWriter) {
	if !valueInList(config.GPIOType, turnOnOffList) {
		writeErrorResponse(w, "GPIO Type cannot be turned off")
		return
	}
	err := drivers.TurnOff(config.GPIOPin)
	if err != nil {
		writeErrorResponse(w, "%s", err)
		return
	} else {
		writeSuccessResponse(w, "Turned off")
	}
}

var mu sync.RWMutex

func (apiServer APIServer) getValue(w http.ResponseWriter) {
	if !valueInList(config.GPIOType, getValueList) {
		writeErrorResponse(w, "Cannot get value for this GPIO type, only options are turnOn or turnOff")
		return
	}

	var response any
	var err error
	var cacheUsed bool = false
	sensor.Type = config.GPIOType

	switch gpioType := sensor.Type; {
	case gpioType == "sht3x":
		mu.Lock() // prevent spamming sensor
		response, cacheUsed, err = sensor.GetValue(config.Unit, config.CacheTTL)
		mu.Unlock()
	}

	if err != nil {
		writeErrorResponse(w, "Error getting value: %v", err)
	} else if config.Debug && cacheUsed {
		log.Println("Values retrieved using cache")
	}

	writeResponse(w, response)

}
