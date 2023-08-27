package presscontinue

import (
	"github.com/anotherhadi/getchar"
)

// Ask a Yes/No question
func PressContinue() (state bool, err error) {
	ascii, _, err := getchar.GetChar()
	if err != nil {
		return
	}
	state = ascii == 13
	return
}
