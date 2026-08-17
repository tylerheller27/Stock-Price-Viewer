package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var finnhubBaseURL = "https://finnhub.io/api/v1" //used in testing

type Quote struct {
	Current            float64 `json:"c"`
	Change             float64 `json:"d"`
	PercentChange      float64 `json:"dp"`
	High               float64 `json:"h"`
	Low                float64 `json:"l"`
	OpenPrice          float64 `json:"o"`
	PreviousClosePrice float64 `json:"pc"`
}

func getQuote(ticker string, apiKey string) (Quote, error) {

	//Values is just a map
	values := url.Values{}
	values.Set("symbol", ticker)
	values.Set("token", apiKey)

	fullURL := finnhubBaseURL + "/quote?" + values.Encode()

	//
	resp, err := http.Get(fullURL)
	if err != nil {
		return Quote{}, fmt.Errorf("finnhub: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Quote{}, fmt.Errorf("finnhub: unexpected status %d", resp.StatusCode)
	}

	//fmt.Println(fullURL) was for testing if url was correct
	fmt.Println(resp)

	// Placeholder return — body reading/JSON parsing is the next step.
	return Quote{}, nil
}
