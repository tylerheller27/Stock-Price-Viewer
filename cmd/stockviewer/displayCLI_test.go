package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestDisplayCLI_workingCase(t *testing.T) {

	//working case of a normal, non-exit ticker
	scanner := bufio.NewScanner(strings.NewReader("voo\n"))

	got, err := displayCLI(scanner)
	want := "VOO"

	if err != nil {
		t.Fatalf("displayCLI() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("displayCLI() = %q, want %q", got, want)
	}
} //func
