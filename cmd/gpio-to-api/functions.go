package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	GPIO_TYPE string
	GPIO_PIN  int
	Error     string
}

type SuccessResponse struct {
	GPIO_TYPE string
	GPIO_PIN  int
	Message   string
}

func writeErrorResponse(w http.ResponseWriter, s string, v ...any) {
	var response ErrorResponse
	response.GPIO_TYPE = config.GPIOType
	response.GPIO_PIN = config.GPIOPin
	response.Error = fmt.Sprintf(s, v...)
	writeResponse(w, response)
}

func writeSuccessResponse(w http.ResponseWriter, s string, v ...any) {
	var response SuccessResponse
	response.GPIO_TYPE = config.GPIOType
	response.GPIO_PIN = config.GPIOPin
	response.Message = fmt.Sprintf(s, v...)
	writeResponse(w, response)
}

// writes api response to webpage and app log
func writeResponse(w http.ResponseWriter, response any) {
	responseString, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error writing API response: %v", err)
		fmt.Fprintf(w, "Error writing API response: %v", err)
		return
	}
	log.Println(string(responseString))
	json.NewEncoder(w).Encode(response)
}

func valueInList(value string, list []string) bool {
	for _, listVal := range list {
		if listVal == value {
			return true
		}
	}
	return false
}
