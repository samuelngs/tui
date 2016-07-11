// +build windows plan9 solaris

package tui

// GetWindowSize returns the window bound
func getWinSize() *Bound {
	ws := new(tmsize)
	ws.Col = 80
	ws.Row = 24
	return &Bound{
		Top:    int(ws.Ypixel),
		Left:   int(ws.Xpixel),
		Bottom: int(ws.Ypixel + ws.Row),
		Right:  int(ws.Xpixel + ws.Col),
		Height: uint(ws.Row),
		Width:  uint(ws.Col),
	}
}
