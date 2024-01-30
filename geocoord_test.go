package geocoord

import (
	"fmt"
	"testing"
)

func TestGetCoordinates(t *testing.T) {
	postalCode := "87109"
	coordinates, err := GetCoordinates(postalCode)

	if err != nil {
		fmt.Printf("Error: %v/n", err)
		return
	}

	fmt.Printf("Latitude: %f, Longitude: %f\n", coordinates.Latitude, coordinates.Longitude)
}
