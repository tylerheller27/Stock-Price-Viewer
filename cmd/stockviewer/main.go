package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Please Enter A Stock Ticker")
	fmt.Println("Please Type Exit to Terminate Program")

	for {
		ticker, err := displayCLI()

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

		fmt.Printf("Ticker: %v\n", ticker)
	}

} //main
