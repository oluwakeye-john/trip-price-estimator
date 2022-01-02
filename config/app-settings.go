package config

type Settings struct {
	BASE_FARE        float64
	PRICE_PER_MINUTE float64
	PRICE_PER_KM     float64
	MINIMUM_FARE     float64
}

var AppSettings Settings

func InitSettings() {
	AppSettings.BASE_FARE = 200
	AppSettings.PRICE_PER_KM = 50
	AppSettings.PRICE_PER_MINUTE = 10
	AppSettings.MINIMUM_FARE = 350
}
