package main

import "fmt"

func printStockInfo(ticker string, quote Quote) {
	fmt.Printf("Ticker: %s\n", ticker)
	fmt.Printf("Current Price: %.2f\n", quote.Current)
	fmt.Printf("Change: %.2f (%.2f%%)\n", quote.Change, quote.PercentChange)
	fmt.Printf("Open: %.2f\n", quote.OpenPrice)
	fmt.Printf("High: %.2f\n", quote.High)
	fmt.Printf("Low: %.2f\n", quote.Low)
	fmt.Printf("Previous Close: %.2f\n", quote.PreviousClosePrice)
} //printStockInfo
