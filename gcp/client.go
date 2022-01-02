package gcp

import (
	"log"
	"net/http"
	"net/url"

	"github.com/oluwakeye-john/trip-price-estimator/config"
)

type MapClient struct {
}

func NewMapClient() MapClient {
	return MapClient{}
}

func (m *MapClient) Request(path string, params url.Values) (*http.Response, error) {
	GCP_API_KEY := config.MustGetEnv("GCP_API_KEY")
	params.Add("key", GCP_API_KEY)

	fmt_url := url.URL{
		Scheme:   "https",
		Host:     "maps.googleapis.com",
		Path:     path,
		RawQuery: params.Encode(),
	}
	log.Println("Request URL: ", fmt_url.String())
	return http.Get(fmt_url.String())
}
