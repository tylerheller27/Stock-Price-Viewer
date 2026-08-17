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
	apiKey, err := loadAPIKey() // loading api key from config.go

	//checking to see if API key is empty and terminating program if itis
	if errors.Is(err, ErrMissingAPIKey) {
		fmt.Println("no API key loaded in env file, terminating program")
		return
	}

	for {

		//calling DisplayCLI function to get user input of a stock ticker
		ticker, err := displayCLI(scanner)

		//looking for user response "exit" to terminate program
		if errors.Is(err, ErrExit) {
			fmt.Println("Goodbye!")
			return
		}

		//sending user input to isValidTicker function to check if thier response
		// is betweeen 1-5 character & all caracters are letters
		// returns true if its a valid ticker
		if !isValidTicker(ticker) {
			fmt.Println("Please Enter A Valid Ticker")
			continue //restarts looping looking for correct ticker
		} //if

		//building the api call with api key
		//program will aleardy terminate if the API key is not properly if the API key is not
		//setup as a OS environment variable. we will send the verified ticker as well as
		//api key to getQuote function to build the request

		getQuote(ticker, apiKey) // this is not fully built out so it should just diplay the url

		//fmt.Println(apiKey) //using apiKey so the program compiles.. delete later

	} //for

} //main
