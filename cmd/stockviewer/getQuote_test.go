package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuote_workingCase(t *testing.T) {

	//fake local server standing in for the real Finnhub API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"c":710.27,"d":-3.34,"dp":-0.468,"h":714.045,"l":710.1,"o":713.5,"pc":713.61}`))
	}))
	defer server.Close()

	//point getQuote at the fake server instead of the real Finnhub URL
	original := finnhubBaseURL
	finnhubBaseURL = server.URL
	defer func() { finnhubBaseURL = original }()

	got, err := getQuote("VOO", "test-key")
	want := Quote{
		Current:            710.27,
		Change:             -3.34,
		PercentChange:      -0.468,
		High:               714.045,
		Low:                710.1,
		OpenPrice:          713.5,
		PreviousClosePrice: 713.61,
	}

	if err != nil {
		t.Fatalf("getQuote() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("getQuote() = %+v, want %+v", got, want)
	}
} //func

func TestGetQuote_badStatus(t *testing.T) {

	//fake server that always fails, standing in for a Finnhub outage/bad key
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	original := finnhubBaseURL
	finnhubBaseURL = server.URL
	defer func() { finnhubBaseURL = original }()

	_, err := getQuote("VOO", "test-key")

	if err == nil {
		t.Fatal("getQuote() error = nil, want non-nil for a non-200 status")
	}
} //func

func TestGetQuote_malformedJSON(t *testing.T) {

	//fake server returning 200 but with a broken JSON body
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"c": not valid json`))
	}))
	defer server.Close()

	original := finnhubBaseURL
	finnhubBaseURL = server.URL
	defer func() { finnhubBaseURL = original }()

	_, err := getQuote("VOO", "test-key")

	if err == nil {
		t.Fatal("getQuote() error = nil, want non-nil for malformed JSON")
	}
} //func

func TestGetQuote_networkFailure(t *testing.T) {

	//no server at all -- port 1 on localhost has nothing listening on it
	original := finnhubBaseURL
	finnhubBaseURL = "http://127.0.0.1:1"
	defer func() { finnhubBaseURL = original }()

	_, err := getQuote("VOO", "test-key")

	if err == nil {
		t.Fatal("getQuote() error = nil, want non-nil for an unreachable server")
	}
} //func
