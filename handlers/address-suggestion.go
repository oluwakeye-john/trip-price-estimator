package handlers

import (
	"github.com/oluwakeye-john/trip-price-estimator/gcp"
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
)

func AddressSuggestion(input string) (*model.AddressSuggestion, error) {
	address_suggestion := &model.AddressSuggestion{}

	result, err := gcp.AutoCompleteRequest(input)
	if err != nil {
		return address_suggestion, err
	}

	for _, i := range result.Predictions {
		address_suggestion.Predictions = append(address_suggestion.Predictions, &model.Prediction{
			Description: i.Description,
		})
	}

	return address_suggestion, nil
}
