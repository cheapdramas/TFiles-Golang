package main

import (
	"fmt"
	"tfiles/terminal"
)

func main() {
	screen := terminal.NewScreen()

	for (true) {
		screen.ScreenUpdateSize()
		fmt.Println(screen.Width)
	}
}
