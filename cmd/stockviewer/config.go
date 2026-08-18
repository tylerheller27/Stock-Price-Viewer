package main

import (
	"errors"
	"os"
)

var ErrMissingAPIKey = errors.New("FINNHUB_API_KEY environment variable is not set")

func loadAPIKey() (string, error) {
	apiKey := os.Getenv("FINNHUB_API_KEY")

	if apiKey == "" {
		return apiKey, ErrMissingAPIKey
	} else {
		return apiKey, nil
	}

} //loadAPIKey
