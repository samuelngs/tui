package tui

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

// terminal window size
type tmsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Bound is the size of the element and the position of the viewport
type Bound struct {
	Top    int
	Left   int
	Bottom int
	Right  int
	Height uint
	Width  uint
}

func (v *Bound) String() string {
	return fmt.Sprintf(`
	top     : %v
	left    : %v
	bottom  : %v
	right   : %v
	height  : %v
	width   : %v
	`, v.Top, v.Left, v.Bottom, v.Right, v.Height, v.Width)
}

// Terminal struct
type Terminal struct {
	buf    *bytes.Buffer
	stdin  *bufio.Writer
	stdout *bufio.Writer
}

// Term creates a new terminal session
func Term() *Terminal {
	return &Terminal{
		buf:    new(bytes.Buffer),
		stdin:  bufio.NewWriter(os.Stdin),
		stdout: bufio.NewWriter(os.Stdout),
	}
}

// Printf to print message to tty
func (v *Terminal) Printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(v.buf, f, args...)
}

// Write byte message to tty
func (v *Terminal) Write(b []byte) (int, error) {
	return v.buf.Write(b)
}

// Flush message
func (v *Terminal) Flush() {
	v.stdout.Write(v.buf.Bytes())
	v.stdout.Flush()
}

// Reset pending print message(s)
func (v *Terminal) Reset() {
	v.buf.Reset()
}

// Clear screen
func (v *Terminal) Clear() {
	v.Printf("\033[2J")
}

// MoveCursor to move cursor to given position
func (v *Terminal) MoveCursor(x int, y int) {
	v.Printf("\033[%d;%dH", x, y)
}

// Size to return the size of the terminal
func (v *Terminal) Size() *Bound {
	return getWinSize()
}

// Render to render component(s)
func (v *Terminal) Render(i IComponent) {
	for {
		tmpl := compile(i)
		v.Clear()
		v.Write(tmpl)
		v.Flush()
		time.Sleep(1 * time.Second)
	}
}
