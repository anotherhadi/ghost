<p align="center">
  <img width="150px" src="https://cdn.nostr.build/i/0d8b39f5ea38c833592cb06ef3c81a1c885334338fcc8bd2561c345774f379a2.png" />
</p>

# Ghost UI
## TUI Tools for Go

The 'ghost' package provides a collection of tools for creating Text-based User Interfaces (TUI). It offers features such as input fields, selections, and other components essential for building interactive text-based applications.

### Features:

- Input field
- Selection
- Checkbox
- Yes/No
- Press to continue
- Print with prefix

---

## Usage

### Installation:

```bash
go get github.com/anotherhadi/ghost
```

Here's how you can use the components of this module in your code:

### Input

```go
package main

import (
  "github.com/anotherhadi/ghost/print"
  "github.com/anotherhadi/ghost/input"
)

func main() {
	print.Input("What's your name ?")
	name := input.Input("John")
	print.Info("Your name is " + name)
}
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/5bababe570e3c7ce0d5c7632dee84f2f07580fb2d8c7d7f325d9d36264ef5ffe.png" />

### Selection

```go
package main

import (
  "github.com/anotherhadi/ansi"
  "github.com/anotherhadi/ghost/print"
  "github.com/anotherhadi/ghost/selection"
)

func main() {
	print.Input("Choose an option:", ansi.BrightBlack+"(Arrow or j/k)")
	options := []string{"Option 1", "Option 2", "Option 3", "Option 4"}
	selected, err := selection.Selection(options)
	if err != nil {
		print.Error("Error:", err)
		return
	}
	print.Info("You choose", options[selected])
}
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/0afb186863c5940e4bc5d15d01a53319c69ec79f23310aa164ac2ac8bc461dda.png" />

### Checkbox

```go
package main

import (
	"github.com/anotherhadi/ansi"
  "github.com/anotherhadi/ghost/print"
  "github.com/anotherhadi/ghost/checkbox"
)

func main() {
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
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/928342487940eb1eaeda385358f95a0563ca2b1fe6ed4b6878e63f1fe8fa0faa.png" />

### Yes/No

```go
package main

import (
  "github.com/anotherhadi/ghost/print"
  "github.com/anotherhadi/ghost/yesno"
)

func main() {
	print.Input("Do you want to add a star to Ghost ?")
	answer := yesno.YesNo(true)
	if answer {
		print.Success("Youhouuu !!")
	} else {
		print.Error("Wrong answer.")
	}
}
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/a0e8075900319ea7ac6ef099da99fe96d293cea96842971e23749b1373ca18c4.png" />

### Press to Continue

```go
package main

import (
  "github.com/anotherhadi/ghost/print"
  "github.com/anotherhadi/ghost/presscontinue"
)

func main() {
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
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/5ae5038a0148ba369b013a5372c4ce574d2fa0e3677106d7212a8f3eb3efcd67.png" />

### Print (with prefix)

```go
package main

import (
  "github.com/anotherhadi/ghost/print"
)

func main() {
	print.Info("Info message")
	print.Input("Input message")
	print.Success("Success !")
	print.Warning("Warning..")
	print.Error("Ouch, this one is an error")
}
```

**Result:**

<img width="50%" src="https://cdn.nostr.build/i/3fb203ef69ff8bbc2b0d998e9f81d2aa3693fe6fa19da033b36852e89aea3d35.png" />

---

Examples available in "test/test.go"

---

<img src="https://img.buymeacoffee.com/button-api/?text=Buy me a cookie&emoji=🍪&slug=anotherhadi&button_colour=eed2cc&font_colour=000000&font_family=Inter&outline_colour=ffffff&coffee_colour=ff0000" />
