package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var finnhubBaseURL = "https://finnhub.io/api/v1" //used in testing

// Finnhub doesn't return an HTTP error for an unknown ticker -- it responds
// with 200 OK and every price field set to 0 instead. ErrTickerNotFound lets
// getQuote surface that as a distinct, checkable error via errors.Is.
var ErrTickerNotFound = errors.New("ticker not found")

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

	// resp.Body is a stream, not usable data yet -- io.ReadAll drains it
	// into memory as raw bytes. This is a separate failure point from the
	// network/status checks above: the connection could succeed and still
	// fail partway through reading (e.g. connection drops mid-transfer).
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, fmt.Errorf("finnhub: reading response failed: %w", err)
	}

	// var quote Quote creates a zero-valued Quote (all fields 0) so
	// &quote has something real to point at. json.Unmarshal needs a
	// pointer because it must modify the value in place -- without the
	// pointer, it would only be writing into a throwaway copy.
	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return Quote{}, fmt.Errorf("finnhub: parsing response failed: %w", err)
	}

	// json.Unmarshal matches JSON keys ("c", "h", etc.) to the struct's
	// `json:"..."` tags. Any JSON field without a matching tag (like
	// Finnhub's timestamp "t") is just silently ignored.

	// Checking multiple fields (not just Current) makes this more robust
	// against a single stray zero -- a real ticker is extremely unlikely
	// to have every one of these fields be exactly 0 at once.
	if quote.Current == 0 && quote.High == 0 && quote.Low == 0 && quote.OpenPrice == 0 && quote.PreviousClosePrice == 0 {
		return Quote{}, ErrTickerNotFound
	}

	return quote, nil
}
