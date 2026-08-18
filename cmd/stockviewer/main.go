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

	apiKey, err := loadAPIKey() // loading api key from env variable via config.go

	//checking to see if API key is empty and terminating program if it is
	if errors.Is(err, ErrMissingAPIKey) {
		fmt.Println("no API key loaded in env file, terminating program")
		return
	}

	//for loop to continually ask for user input; the only way it exits is
	//the user typing "exit" (checked below via ErrExit). The missing-API-key
	//case above already returned before we ever get here.

	for {

		//calling displayCLI function to get user input of a stock ticker
		ticker, err := displayCLI(scanner)

		//looking for user response "exit" to terminate program
		if errors.Is(err, ErrExit) {
			fmt.Println("Goodbye!")
			return
		}

		//sending user input to isValidTicker function to check if their response
		// is between 1-5 characters & all characters are letters
		// returns true if it's a valid ticker

		if !isValidTicker(ticker) {
			fmt.Println("Please Enter A Valid Ticker")
			continue //restarts looping looking for correct ticker
		} //if

		//ticker is valid, apiKey was already confirmed present at startup --
		//send both to getQuote to build the request and fetch live data.
		//Note: the key's *validity* isn't checked until Finnhub responds --
		//an invalid key would surface here as an "unexpected status" error.
		quote, err := getQuote(ticker, apiKey)
		if err != nil {
			//ErrTickerNotFound is Finnhub's "200 OK with all-zero fields"
			//quirk for an unknown symbol -- everything else (network/status/
			//parse failures) is a real error, shown as-is.
			if errors.Is(err, ErrTickerNotFound) {
				fmt.Println("Ticker not found")
			} else {
				fmt.Println("Error fetching quote:", err)
			}
			continue
		}

		printStockInfo(ticker, quote)

	} //for

} //main
