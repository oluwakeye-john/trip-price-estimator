package services

import (
	"math"

	"github.com/oluwakeye-john/trip-price-estimator/config"
)

type Trip struct {
	Distance float64
	Duration float64
	Settings config.Settings
}

func (t *Trip) ComputeTripFee() float64 {
	base_fare := t.Settings.BASE_FARE
	price_per_km := t.Settings.PRICE_PER_KM
	price_per_minute := t.Settings.PRICE_PER_MINUTE
	minimum_fare := t.Settings.MINIMUM_FARE

	price := base_fare + ((t.Distance / 1000) * price_per_km) + ((t.Duration / 60) * price_per_minute)

	if price < minimum_fare {
		price = minimum_fare
	}

	// round price to nearest N10
	return math.Round(price/10) * 10
}
