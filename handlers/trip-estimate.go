package handlers

import (
	"errors"

	"github.com/oluwakeye-john/trip-price-estimator/gcp"
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
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

	trip_estimate.Price = 2000
	trip_estimate.Distance = element.Distance.Value
	trip_estimate.Duration = element.Duration.Value

	return trip_estimate, nil
}
