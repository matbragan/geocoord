package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var logFormat = "%s - %s - %s\n"

func init() {
	// Create or open the log file for writing
	logFile, err := os.OpenFile("coordinates_go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Set log output to the log file
	log.SetOutput(logFile)

	// Set log flags to include timestamp
	log.SetFlags(log.LstdFlags)
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getCoordinatesFromAPI(cep, street, neighborhood, city, state string) (*Coordinates, error) {
	apiURL := "http://127.0.0.1:5000/get_coordinates"

	// Prepare JSON data for the API request
	requestData := map[string]string{
		"cep":          cep,
		"street":       street,
		"neighborhood": neighborhood,
		"city":         city,
		"state":        state,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	// Make a POST request to the API
	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response JSON
	var coordinates Coordinates
	err = json.Unmarshal(body, &coordinates)
	if err != nil {
		return nil, err
	}

	return &coordinates, nil
}
