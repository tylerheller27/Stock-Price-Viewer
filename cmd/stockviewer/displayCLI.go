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

	exitString := strings.ToUpper(userInput)

	//checking to see if user typed "Exit" so we can termineate the program.
	if exitString == "EXIT" {
		return "", ErrExit
	}

	return userInput, nil

} //displayCLI
