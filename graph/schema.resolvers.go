package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oluwakeye-john/trip-price-estimator/graph/generated"
	"github.com/oluwakeye-john/trip-price-estimator/graph/model"
	"github.com/oluwakeye-john/trip-price-estimator/handlers"
)

func (r *mutationResolver) RequestTrip(ctx context.Context, email string, origin string, destination string) (*model.RequestTrip, error) {
	return handlers.RequestRide(origin, destination, email)
}

func (r *queryResolver) Geocoding(ctx context.Context, input string) (*model.Geocoding, error) {
	return handlers.Geocoding(input)
}

func (r *queryResolver) ReverseGeocoding(ctx context.Context, latitude float64, longitude float64) (*model.ReverseGeocoding, error) {
	return handlers.ReverseGeocoding(latitude, longitude)
}

func (r *queryResolver) AddressSuggestion(ctx context.Context, input string) (*model.AddressSuggestion, error) {
	return handlers.AddressSuggestion(input)
}

func (r *queryResolver) TripEstimate(ctx context.Context, origin string, destination string) (*model.TripEstimate, error) {
	return handlers.TripEstimate(origin, destination)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
