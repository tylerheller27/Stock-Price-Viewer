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

func TestDisplayCLI_endOfFile(t *testing.T) {

	//True end of file.  no input at all, not even a blank line
	scanner := bufio.NewScanner(strings.NewReader(""))

	//we only care about what err displayCLI returns so we just discard the ticket with _
	_, err := displayCLI(scanner)
	want := ErrExit

	if err != want {
		t.Fatalf("displayCLI() error = %v, want ErrExit", err)
	}

} //func

func TestDisplayCLI_emptyLine(t *testing.T) {

	//Not EOF -- the user pressed Enter with nothing typed. displayCLI
	//doesn't reject this itself; it just hands back an empty string.

	scanner := bufio.NewScanner(strings.NewReader("\n"))

	got, err := displayCLI(scanner)
	want := ""

	if err != nil {
		t.Fatalf("displayCLI() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("displayCLI() = %q, want %q", got, want)
	}
} //func

func TestDisplayCLI_whitespaceOnly(t *testing.T) {

	//displayCLI doesn't trim whitespace -- it should come back unchanged,
	//since content-validation is isValidTicker's job, not displayCLI's.

	scanner := bufio.NewScanner(strings.NewReader("    \n"))

	got, err := displayCLI(scanner)
	want := "    "

	if err != nil {
		t.Fatalf("displayCLI() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("displayCLI() = %q, want %q", got, want)
	}
} //func
