package print

import (
	"fmt"
	"github.com/anotherhadi/ansi"
)

func Input(a ...any) {
	fmt.Print(ansi.BrightBlack, "[", ansi.Green, "?", ansi.BrightBlack, "] ", ansi.Reset, ansi.Bold)
	fmt.Println(a...)
	fmt.Print(ansi.Reset)
}

func Error(a ...any) {
	fmt.Print(ansi.BrightBlack, "[", ansi.Red, "x", ansi.BrightBlack, "] ", ansi.Reset, ansi.Bold)
	fmt.Println(a...)
	fmt.Print(ansi.Reset)
}

func Warning(a ...any) {
	fmt.Print(ansi.BrightBlack, "[", ansi.BrightRed, "~", ansi.BrightBlack, "] ", ansi.Reset, ansi.Bold)
	fmt.Println(a...)
	fmt.Print(ansi.Reset)
}

func Success(a ...any) {
	fmt.Print(ansi.BrightBlack, "[", ansi.Green, "v", ansi.BrightBlack, "] ", ansi.Reset, ansi.Bold)
	fmt.Println(a...)
	fmt.Print(ansi.Reset)
}

func Info(a ...any) {
	fmt.Print(ansi.BrightBlack, "[-] ", ansi.Reset, ansi.Bold)
	fmt.Println(a...)
	fmt.Print(ansi.Reset)
}
