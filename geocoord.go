package geocoord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// APIEndpoint is the URL of your Python API.
var APIEndpoint = "https://geo-coord-api.vercel.app/"

// Coordinates represents the latitude and longitude.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetCoordinates retrieves the coordinates for a given ZIP code using the Python API.
func GetCoordinates(zipCode string) (Coordinates, error) {

	// Define the body data
	requestBody, err := json.Marshal(map[string]string{"zip_code": zipCode})
	if err != nil {
		return Coordinates{}, fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Make a POST request
	response, err := http.Post(APIEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return Coordinates{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Coordinates{}, fmt.Errorf("unexpected response status: %v", response.Status)
	}

	var coordinates Coordinates
	if err := json.NewDecoder(response.Body).Decode(&coordinates); err != nil {
		return Coordinates{}, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return coordinates, nil
}
