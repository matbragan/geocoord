package geocoord

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestgetCoordinates(t *testing.T) {
	// Mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := "/"
		if r.URL.String() != expectedURL {
			t.Errorf("Expected URL: %s, got: %s", expectedURL, r.URL.String())
		}

		expectedZIP := "12345"
		body := make(map[string]string)
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if body["zip_code"] != expectedZIP {
			t.Errorf("Expected ZIP code: %s, got: %s", expectedZIP, body["zip_code"])
		}

		// Return mock response
		response := `{"latitude": 40.7128, "longitude": -74.0060}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	// Set APIEndpoint to mock server URL
	APIEndpoint = ts.URL

	// Test GetCoordinatesByZIP
	coordinates, err := getCoordinates("12345")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedLatitude := 40.7128
	expectedLongitude := -74.0060
	if coordinates.Latitude != expectedLatitude || coordinates.Longitude != expectedLongitude {
		t.Errorf("Expected latitude: %f, longitude: %f, got latitude: %f, longitude: %f",
			expectedLatitude, expectedLongitude, coordinates.Latitude, coordinates.Longitude)
	}
}
