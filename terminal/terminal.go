package terminal

import (
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


func enableRawMode(fd int) (*syscall.Termios, error) {
    var orig syscall.Termios // saved current settings
    if _, _, err := syscall.Syscall6( // get terminal settings
        syscall.SYS_IOCTL,
        uintptr(fd),
        uintptr(syscall.TCGETS),
        uintptr(unsafe.Pointer(&orig)),
        0, 0, 0,
    ); err != 0 {
        return nil, err
    }

    raw := orig

    // отключаем canonical mode и echo
    raw.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.ISIG | syscall.IEXTEN
    raw.Cc[syscall.VMIN] = 1  // минимальное количество байт для read
    raw.Cc[syscall.VTIME] = 0 // таймаут для read

    if _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        uintptr(fd),
        uintptr(syscall.TCSETS),
        uintptr(unsafe.Pointer(&raw)),
        0, 0, 0,
    ); err != 0 {
        return nil, err
    }

    return &orig, nil
}

func disableRawMode(fd int, orig *syscall.Termios) error {
    if _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        uintptr(fd),
        uintptr(syscall.TCSETS),
        uintptr(unsafe.Pointer(orig)),
        0, 0, 0,
    ); err != 0 {
        return err
    }
    return nil
}



func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Prepare() {
	
}

//method for Screen
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
