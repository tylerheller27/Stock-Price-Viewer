package main

import (
	"fmt"
	"net/url"
)

var finnhubBaseURL = "https://finnhub.io/api/v1"

type Quote struct {
	Current            float64 `json:"c"`
	Change             float64 `json:"d"`
	PercentChange      float64 `json:"dp"`
	High               float64 `json:"h"`
	Low                float64 `json:"l"`
	OpenPrice          float64 `json:"o"`
	PreviousClosePrice float64 `json:"pc"`
}

func getQuote(ticker string, apiKey string) {

	values := url.Values{}
	values.Set("symbol", ticker)
	values.Set("token", apiKey)

	fullURL := finnhubBaseURL + "/quote?" + values.Encode()

	fmt.Println(fullURL)
}
