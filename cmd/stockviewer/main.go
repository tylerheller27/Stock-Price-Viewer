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

	// loading api key from config.go
	apiKey, err := loadAPIKey() // loading api key from env variable from config.go

	//checking to see if API key is empty and terminating program if it is
	if errors.Is(err, ErrMissingAPIKey) {
		fmt.Println("no API key loaded in env file, terminating program")
		return
	}

	//for loop to continually ask for user input unless an error terminates the program
	// ex: user types in exit or apiKey env is empty

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

		//building the api call with api key
		//program will terminate if the API key FINNHUB_API_KEY is empty
		//api key validity is not checked in this step.
		//sends valid ticker and api key to getQuote function to build the request
		//and send the request out

		quote, err := getQuote(ticker, apiKey)
		if err != nil {
			fmt.Println("Error fetching quote:", err)
			continue
		}

		printStockInfo(ticker, quote)

	} //for

} //main
