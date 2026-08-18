package main

import (
	"bufio"
	"errors"
	"strings"
)

var ErrExit = errors.New("user requested exit")

func displayCLI(scanner *bufio.Scanner) (string, error) {

	// scanner.Scan() returns false when there's no more input (EOF) — e.g.
	// the user hits Ctrl+D, or piped input runs out. This used to go
	// unchecked, which caused an infinite loop: once Scan() starts failing,
	// Text() silently keeps returning "", which never matches "exit" below,
	// so the program never terminated. Discovered during manual testing
	// with piped multi-line input, which produced 155 million lines of
	// output before being killed.
	if !scanner.Scan() {
		return "", ErrExit
	}
	userInput := scanner.Text()

	// Uppercasing here isn't just for the "exit" comparison below -- it's
	// also what gets returned as the ticker. Verified directly against the
	// live API: Finnhub's /quote endpoint is case-sensitive and silently
	// returns all-zero fields for a lowercase symbol (e.g. "voo") while a
	// real quote comes back for uppercase ("VOO"). Normalizing case here,
	// once, means every ticker downstream is guaranteed usable.
	upperCaseInput := strings.ToUpper(userInput)

	// Use a sentinel error to represent the intentional "EXIT" control-flow
	// path. The caller can detect this with errors.Is(err, ErrExit) and
	// stop the loop without treating it as regular input.

	if upperCaseInput == "EXIT" {
		return "", ErrExit
	}

	return upperCaseInput, nil

} //displayCLI
