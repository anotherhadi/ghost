package input

import (
	"fmt"
	"github.com/anotherhadi/ansi"
)

// Ask user for input with a default return
func Input(defaultState string) string {
	var input string

	fmt.Print(ansi.BrightBlack)
	if defaultState != "" {
		fmt.Print("(", defaultState, ") ")
	}

	fmt.Print("> ", ansi.Reset)

	fmt.Scanln(&input)
	fmt.Print(ansi.Reset)

	if input == "" {
		return defaultState
	}
	return input
}
