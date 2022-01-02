package gcp

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type textValue struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

type DistanceMatrixResult struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Status               string   `json:"status"`
	Rows                 []struct {
		Elements []struct {
			Distance textValue `json:"distance"`
			Duration textValue `json:"duration"`
			Status   string    `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
}

func DistanceMatrixRequest(origins string, destinations string) (DistanceMatrixResult, error) {
	params := url.Values{}

	params.Add("origins", origins)
	params.Add("destinations", destinations)

	map_client := MapClient{}

	r, err := map_client.Request("/maps/api/distancematrix/json", params)

	distance_matrix_result := DistanceMatrixResult{}

	if err != nil {
		return distance_matrix_result, err
	}

	if err := json.NewDecoder(r.Body).Decode(&distance_matrix_result); err != nil {
		return distance_matrix_result, err
	}

	if distance_matrix_result.Status != "OK" {
		return distance_matrix_result, fmt.Errorf("%s", distance_matrix_result.Status)
	}

	return distance_matrix_result, nil
}
