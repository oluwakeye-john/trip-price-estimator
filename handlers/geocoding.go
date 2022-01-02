package handlers

import (
	"errors"

	"github.com/oluwakeye-john/trip-price-estimator/gcp"
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
)

func ReverseGeocoding(latitude float64, longitude float64) (*model.ReverseGeocoding, error) {
	reverse_geocoding := &model.ReverseGeocoding{}

	result, err := gcp.ReverseGeocodingRequest(latitude, longitude)

	if err != nil {
		return reverse_geocoding, err
	}

	if len(result.Results) < 1 {
		return reverse_geocoding, errors.New("no result")
	}

	reverse_geocoding.Description = result.Results[0].FormattedAddress

	return reverse_geocoding, nil
}

func Geocoding(input string) (*model.Geocoding, error) {
	geocoding := &model.Geocoding{}

	result, err := gcp.GeocodingRequest(input)

	if err != nil {
		return geocoding, err
	}

	if len(result.Results) < 1 {
		return geocoding, errors.New("no result")
	}

	geocoding.Latitude = result.Results[0].Geometry.Location.Lat
	geocoding.Longitude = result.Results[0].Geometry.Location.Lng

	return geocoding, nil
}
