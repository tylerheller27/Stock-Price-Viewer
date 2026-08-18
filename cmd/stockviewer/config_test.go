package main

import (
	"errors"
	"testing"
)

func TestLoadAPIKey_workingCase(t *testing.T) {

	t.Setenv("FINNHUB_API_KEY", "test-api-key")

	got, err := loadAPIKey()
	want := "test-api-key"

	if err != nil {
		t.Fatalf("loadAPIKey() error = %v, want nil", err)
	}

	if got != want {
		t.Errorf("loadAPIKey() = %q, want %q", got, want)
	}
}

func TestLoadAPIKey_nilCase(t *testing.T) {

	t.Setenv("FINNHUB_API_KEY", "")

	got, err := loadAPIKey()
	want := ErrMissingAPIKey

	if !errors.Is(err, want) {
		t.Fatalf("loadAPIKey() error = %v, want %v", err, want)
	}

	if got != "" {
		t.Errorf("loadAPIKey() = %q, want empty string", got)
	}
}
