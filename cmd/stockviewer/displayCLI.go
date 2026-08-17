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

	upperCaseInput := strings.ToUpper(userInput)

	// Use a sentinel error to represent the intentional "EXIT" control-flow
	// path. The caller can detect this with errors.Is(err, ErrExit) and
	// stop the loop without treating it as regular input.

	if upperCaseInput == "EXIT" {
		return "", ErrExit
	}

	return upperCaseInput, nil //

} //displayCLI
