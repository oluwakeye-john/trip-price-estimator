package handlers

import (
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
	"github.com/oluwakeye-john/trip-price-estimator/utils"
)

func RequestRide(origin string, destination string, email string) (*model.RequestTrip, error) {
	estimate, err := TripEstimate(origin, destination)

	request_trip := &model.RequestTrip{}

	if err != nil {
		return request_trip, err
	}

	request_trip.Amount = estimate.Price * 100
	request_trip.Reference = utils.GenerateReference()

	return request_trip, nil
}
