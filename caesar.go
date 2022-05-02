package main

import (
	"fmt"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz"

func caesar(textIn string, method string) (textOut string) {
	var shift int

	fmt.Print("\nPlease enter a shift: ")
	fmt.Scanf("%d\n", &shift)

	if method == "d" { // decrypt (use negative shift)
		shift *= -1
	}

	for i, _ := range textIn {
		var nextIndex int = (i + shift) % len(ALPHABET)
		var nextVal string = ALPHABET[nextIndex:nextIndex+1]
		textOut += nextVal
	}

	return
}
