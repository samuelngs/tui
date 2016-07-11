// +build !windows,!plan9,!solaris

package tui

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// GetWindowSize returns the window bound
func getWinSize() *Bound {
	var TIOCGWINSZ int64
	ws := new(tmsize)
	switch runtime.GOOS {
	case "linux":
		TIOCGWINSZ = syscall.TIOCGWINSZ
	case "darwin":
		TIOCGWINSZ = 1074295912
	}
	r, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if int(r) == -1 {
		panic(os.NewSyscallError("GetWindowSize", err))
	}
	return &Bound{
		Top:    int(ws.Ypixel),
		Left:   int(ws.Xpixel),
		Bottom: int(ws.Ypixel + ws.Row),
		Right:  int(ws.Xpixel + ws.Col),
		Height: uint(ws.Row),
		Width:  uint(ws.Col),
	}
}
