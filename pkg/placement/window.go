package placement

type Window struct {
	Height, Width    int
	XOffset, YOffset int
	Addr             int64
	Father           *Window
}

func NewWindow(height, width, xOffset, yoffset int, addr int64, father *Window) *Window {
	return &Window{
		Height:  height,
		Width:   width,
		XOffset: xOffset,
		YOffset: yoffset,
		Addr:    addr,
		Father:  father,
	}
}
