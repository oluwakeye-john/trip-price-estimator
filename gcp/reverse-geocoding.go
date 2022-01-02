package gcp

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type ReverseGeocodingResult struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	Results      []struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
}

func ReverseGeocodingRequest(lat float64, lng float64) (ReverseGeocodingResult, error) {
	params := url.Values{}
	params.Add("latlng", fmt.Sprintf("%f,%f", lat, lng))

	reverse_geocoding_result := ReverseGeocodingResult{}

	map_client := NewMapClient()
	r, err := map_client.Request("/maps/api/geocode/json", params)

	if err != nil {
		return reverse_geocoding_result, err
	}

	if err := json.NewDecoder(r.Body).Decode(&reverse_geocoding_result); err != nil {
		return reverse_geocoding_result, err
	}

	if reverse_geocoding_result.Status != "OK" {
		return reverse_geocoding_result, fmt.Errorf("%s: %s", reverse_geocoding_result.Status, reverse_geocoding_result.ErrorMessage)
	}

	return reverse_geocoding_result, nil
}
