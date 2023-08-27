package main

import (
	"fmt"

	"github.com/anotherhadi/ansi"
	"github.com/anotherhadi/ghost/checkbox"
	"github.com/anotherhadi/ghost/input"
	"github.com/anotherhadi/ghost/presscontinue"
	"github.com/anotherhadi/ghost/print"
	"github.com/anotherhadi/ghost/selection"
	"github.com/anotherhadi/ghost/yesno"
)

func inputTest() {
	print.Input("What's your name ?")
	name := input.Input("John")
	print.Info("Your name is " + name)
}

func selectionTest() {
	print.Input("Choose an option:", ansi.BrightBlack+"(Arrow or j/k)")
	options := []string{"Option 1", "Option 2", "Option 3", "Option 4"}
	selected, err := selection.Selection(options)
	if err != nil {
		print.Error("Error:", err)
		return
	}
	print.Info("You choose", options[selected])
}

func checkboxTest() {
	print.Input("Choose options:", ansi.BrightBlack+"(Arrow or j/k to move, space to select)")
	options := []string{"Option 1", "Option 2", "Option 3", "Option 4"}
	selected, err := checkbox.Checkbox(options)
	if err != nil {
		print.Error("Error:", err)
		return
	}
	print.Info("You choose:")
	for index, state := range selected {
		if state {
			fmt.Println(ansi.BrightBlack, "- ", ansi.Reset, options[index])
		}
	}
}

func yesNoTest() {
	print.Input("Do you want to add a star to Ghost ?")
	answer := yesno.YesNo(true)
	if answer {
		print.Success("Youhouuu !!")
	} else {
		print.Error("Wrong answer.")
	}
}

func pressContinueTest() {
	print.Input("Press enter to continue, other key to abort")
	answer, err := presscontinue.PressContinue()
	if err != nil {
		print.Error("Error:", err)
		return
	}
	if answer {
		print.Success("Let's continue")
	} else {
		print.Error("Abort the mission")
	}
}

func printTest() {
	print.Info("Info message")
	print.Input("Input message")
	print.Success("Success !")
	print.Warning("Warning..")
	print.Error("Ouch, this one is an error")
}

func main() {
	print.Info("Ghost TUI test")
	options := []string{"Input", "Selection", "Checkbox", "Yes/No", "Press to continue", "Print"}
	print.Input("What do you want to test ?")

	selected, _ := selection.Selection(options)
	switch selected {
	case 0:
		inputTest()
	case 1:
		selectionTest()
	case 2:
		checkboxTest()
	case 3:
		yesNoTest()
	case 4:
		pressContinueTest()
	case 5:
		printTest()
	}

}
