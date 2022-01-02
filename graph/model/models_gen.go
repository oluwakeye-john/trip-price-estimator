// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddressSuggestion struct {
	Predictions []*Prediction `json:"predictions"`
}

type Geocoding struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Prediction struct {
	Description string `json:"description"`
}

type ReverseGeocoding struct {
	Description string `json:"description"`
}

type TripEstimate struct {
	// Estimated price of trip in Naira
	Price float64 `json:"price"`
	// Trip distance in Km
	Distance float64 `json:"distance"`
	// Trip duration in seconds
	Duration float64 `json:"duration"`
}
