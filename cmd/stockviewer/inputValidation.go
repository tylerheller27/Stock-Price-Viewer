package main

import (
	"unicode"
	"unicode/utf8"
)

func isValidTicker(input string) bool {

	for _, r := range input {
		if !unicode.IsLetter(r) {
			return false
		} //if
	} //for

	count := utf8.RuneCountInString(input)
	if count > 5 || count < 1 {
		return false
	}

	return true

} //isValidTicker
