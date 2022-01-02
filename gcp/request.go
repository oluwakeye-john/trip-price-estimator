package gcp

import (
	"log"
	"net/http"
	"net/url"

	"github.com/oluwakeye-john/trip-price-estimator/config"
)

func request(host string, path string, params url.Values) (*http.Response, error) {
	GCP_API_KEY := config.MustGetEnv("GCP_API_KEY")
	params.Add("key", GCP_API_KEY)

	fmt_url := url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     path,
		RawQuery: params.Encode(),
	}
	log.Println("URL: ", fmt_url.String())
	return http.Get(fmt_url.String())
}

func newMapRequest(path string, params url.Values) (*http.Response, error) {
	return request("maps.googleapis.com", path, params)
}
