package main

import (
	"fmt"
	"tfiles/terminal"
)


func input_handler(buf []byte, n int) int {
	if n > 0 {

		if cmpBytes(buf[:n], []byte{'q'}) {
			return 1
		} // QUIT

		if cmpBytes(buf[:n], terminal.KEY_arrow_down) {
			fmt.Println("Arrow down pressed")
		}

		// fmt.Printf("Captured: %v bytes: %v \n", string(buf[:n]), buf[:n])
	}


	return 0
}


func main() {
	terminal.Clear()
	screen := terminal.NewScreen()
	old_state, err := screen.Prepare()
	if err != nil {
		fmt.Println("Failed to prepare terminal screen!")
		panic(err)
	}
	defer terminal.DisableRawMode(old_state) // do when main function ends 


	for {
		screen.ScreenUpdateSize()
		buf, n := terminal.ReadInput()

		exit := input_handler(buf, n)
		
		if exit != 0 {
			return
		}
		

		fmt.Println("hello")

	}
}
