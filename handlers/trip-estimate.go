package handlers

import (
	"errors"

	"github.com/oluwakeye-john/trip-price-estimator/config"
	"github.com/oluwakeye-john/trip-price-estimator/gcp"
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
	"github.com/oluwakeye-john/trip-price-estimator/services"
)

func TripEstimate(origin string, destination string) (*model.TripEstimate, error) {
	res, err := gcp.DistanceMatrixRequest(origin, destination)

	trip_estimate := &model.TripEstimate{}

	if err != nil {
		return trip_estimate, err
	}

	if len(res.Rows) < 1 || len(res.Rows[0].Elements) < 1 {
		return trip_estimate, errors.New("no result")
	}

	element := res.Rows[0].Elements[0]

	trip := services.Trip{
		Distance: element.Distance.Value,
		Duration: element.Duration.Value,
		Settings: config.AppSettings,
	}

	trip_estimate.Distance = trip.Distance
	trip_estimate.Duration = trip.Duration
	trip_estimate.Price = trip.ComputeTripFee()

	return trip_estimate, nil
}
