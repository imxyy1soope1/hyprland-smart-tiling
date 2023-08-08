package types

type Window struct {
	Height, Width    int
	XOffset, YOffset int
	Father           *Window
}
