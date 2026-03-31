package terminal

import (
	"fmt"
	"os"
)

var (
	KEY_ESC = byte(27) 
	KEY_arrow_up = []byte{27, 91, 65}
	KEY_arrow_down = []byte{27, 91, 66}

)

func ReadInput() ([]byte, int) {
	// we assume that raw mode is on
	buf := make([]byte, 3)

	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Printf("Error reading key: %v\n", err)
		return buf, -1
	}
	return buf, n
}
