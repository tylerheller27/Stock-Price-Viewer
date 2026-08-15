package main

import (
	"testing"
)

func TestInputValidation(t *testing.T) {
	got := isValidTicker("voo")
	want := true

	if got != want {
		t.Errorf("isValidTicker(\"voo\") = %v, want %v", got, want)
	}
} //TestInputValidation

func TestInputValidation_EmptyString(t *testing.T) {
	got := isValidTicker("")
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
func TestInputValidation_FivePlusCharacters(t *testing.T) {
	got := isValidTicker("abcdef")
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputValidation_Numbers(t *testing.T) {
	got := isValidTicker("123")
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputValidation_SpecialCharacters(t *testing.T) {
	got := isValidTicker("@%&")
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputValidation_UpperBoundaryCharacters(t *testing.T) {
	got := isValidTicker("fxaix")
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputValidation_LowerBoundaryCharacters(t *testing.T) {
	got := isValidTicker("o")
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputValidation_whitespace(t *testing.T) {
	got := isValidTicker("fx ix")
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
