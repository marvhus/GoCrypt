package main

import (
	"fmt"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz"

func caesar(textIn string, method string) (textOut string) {
	var shift int

	fmt.Print("\nPlease enter a shift: ")
	_, err := fmt.Scanf("%d", &shift)
	// If there is an extra \n, then scn again for it, and discard it.
	if err != nil {
		var discard string
		fmt.Scanln(&discard)
	}

	if method == "d" { // decrypt (use negative shift)
		shift *= -1
	}

	for i, _ := range textIn {
		var nextIndex int = (i + shift) % len(ALPHABET)
		var nextVal string = ALPHABET[nextIndex:nextIndex+1]
		textOut += nextVal
		fmt.Print(nextVal)
	}
	fmt.Println()

	return
}
