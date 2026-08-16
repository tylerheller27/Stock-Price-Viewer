package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Please Enter A Stock Ticker")
	fmt.Println("Please Type Exit to Terminate Program")

	// Created once here and reused across every loop iteration (passed into
	// displayCLI as a parameter), not recreated per call. A fresh
	// bufio.Scanner on os.Stdin each call reads ahead and buffers part of
	// the *next* line internally; recreating it discarded that buffered
	// data, silently swallowing input.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		ticker, err := displayCLI(scanner)

		if errors.Is(err, ErrExit) {
			fmt.Println("Goodbye!")
			return
		}

		if isValidTicker(ticker) {
			fmt.Println("Valid ticker")
			continue
		} else {
			fmt.Println("Please Enter A Valid ticker")
		}

		//fmt.Printf("Ticker: %v\n", ticker)

		testapiKey, _ := loadAPIKey()
		getQuote(ticker, testapiKey)

	}

} //main
