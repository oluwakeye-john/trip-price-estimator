package gcp

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type GeocodingResult struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	Results      []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func GeocodingRequest(address string) (GeocodingResult, error) {
	params := url.Values{}
	params.Add("address", address)

	geocoding_result := GeocodingResult{}

	map_client := NewMapClient()

	r, err := map_client.Request("/maps/api/geocode/json", params)

	if err != nil {
		return geocoding_result, err
	}

	if err := json.NewDecoder(r.Body).Decode(&geocoding_result); err != nil {
		return geocoding_result, err
	}

	if geocoding_result.Status != "OK" {
		return geocoding_result, fmt.Errorf("%s: %s", geocoding_result.Status, geocoding_result.ErrorMessage)
	}

	return geocoding_result, nil
}
