package selection

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/ansi"
	"github.com/anotherhadi/getchar"
)

func printOptions(options []string, cursor int) {
	ansi.CursorRestore()
	for index, option := range options {
		if index == cursor {
			fmt.Print(ansi.Green, "> ", ansi.Reset)
		} else {
			fmt.Print(ansi.BrightBlack, "  ", ansi.Reset)
		}
		fmt.Println(option)
	}
}

func input() (key string, err error) {
	ascii, arrow, err := getchar.GetChar()
	if err != nil {
		return
	}

	key = ""

	if ascii == 107 || arrow == "up" {
		key = "up"
	} else if ascii == 106 || arrow == "down" {
		key = "down"
	} else if ascii == 13 {
		key = "enter"
	} else if ascii == 27 || ascii == 3 || ascii == 101 || ascii == 113 {
		key = "exit"
	}
	return
}

func Selection(options []string) (selected int, err error) {
	ansi.CursorSave()
	ansi.CursorInvisible()
	var cursor int = 0
	var key string
	for {
		printOptions(options, cursor)
		key, err = input()
		if err != nil {
			return
		}
		if key == "up" {
			if cursor != 0 {
				cursor -= 1
			}
		} else if key == "down" {
			if cursor != len(options)-1 {
				cursor += 1
			}
		} else if key == "enter" {
			ansi.CursorRestore()
			ansi.ClearScreenDown()
			selected = cursor
			break
		} else if key == "exit" {
			ansi.CursorVisible()
			err = errors.New("Exited")
			return
		}
	}
	ansi.CursorVisible()
	return
}
