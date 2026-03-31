package terminal

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Screen struct {
	Width uint16
	Height uint16
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func NewScreen() *Screen {
	return &Screen{}
}

func enableRawMode() (*syscall.Termios, error) {
	fd := int(syscall.Stdin)
	var oldState syscall.Termios

	// Get the current terminal attributes
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&oldState)))
	if errno != 0 {
		return nil, errno
	}

	// Modify the attributes to enable raw mode
	newState := oldState
	// Disable canonical mode (ICANON) and echo (ECHO)
	newState.Lflag &^= syscall.ICANON | syscall.ECHO

	// Set the new terminal attributes
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&newState)))
	if errno != 0 {
		return nil, errno
	}

	return &oldState, nil
}

func DisableRawMode(oldState *syscall.Termios) error {
	fd := int(syscall.Stdin)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(oldState)))
	if errno != 0 {
		return errno
	}
	return nil
}

func (s *Screen) Prepare()(*syscall.Termios, error) {
	orig_termios, err := enableRawMode()
	if err != nil {
		fmt.Println("Failed to enable raw mode! ")
	}
	return orig_termios, err
}


func (s *Screen) ScreenUpdateSize() { 
	ws := &winsize{}	

	_, _, err := syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(1), // stdout fd
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)),
    )

	if err != 0 {
		panic(err)
	}

	s.Width = ws.Col
	s.Height = ws.Row
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
