package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrExit = errors.New("user requested exit")

func displayCLI() (string, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	exitString := strings.ToLower(userInput)

	// Use a sentinel error to represent the intentional "EXIT" control-flow
	// path. The caller can detect this with errors.Is(err, ErrExit) and
	// stop the loop without treating it as regular input.

	if exitString == "exit" {
		return "", ErrExit
	}

	return userInput, nil

} //displayCLI
