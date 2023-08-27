package yesno

import (
	"fmt"
	"github.com/anotherhadi/ansi"
	"strings"
)

// Ask a Yes/No question
func YesNo(defaultState bool) bool {
	var input string

	if defaultState {
		fmt.Print(ansi.BrightBlack, "(Y/n) ", ansi.Reset)
	} else {
		fmt.Print(ansi.BrightBlack, "(y/N) ", ansi.Reset)
	}

	fmt.Scanln(&input)

	input = strings.ToLower(input)
	if len(input) == 0 {
		return defaultState
	}

	if input == "y" || input == "yes" {
		return true
	}
	return false
}
