package gcp

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type AutoCompleteResult struct {
	Predictions []struct {
		Description string
	} `json:"predictions"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func AutoCompleteRequest(input string) (AutoCompleteResult, error) {
	params := url.Values{}
	params.Add("input", input)
	params.Add("components", "country:ng")

	var err error
	auto_complete_result := AutoCompleteResult{}

	r, err := newMapRequest("/maps/api/place/autocomplete/json", params)

	if err != nil {
		return auto_complete_result, err
	}

	if err := json.NewDecoder(r.Body).Decode(&auto_complete_result); err != nil {
		return auto_complete_result, err
	}

	if err != nil {
		return auto_complete_result, err
	}

	if auto_complete_result.Status != "OK" {
		return auto_complete_result, fmt.Errorf("%s: %s", auto_complete_result.Status, auto_complete_result.ErrorMessage)
	}

	return auto_complete_result, nil
}
