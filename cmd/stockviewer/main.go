package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		ticker, err := displayCLI()
		if errors.Is(err, ErrExit) {
			fmt.Println("Goodbye!")
			return
		}
		fmt.Println(ticker)
	}

} //main
