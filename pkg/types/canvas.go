package types

type Canvas struct {
	Height, Width    int
	XOffset, YOffset int
	Windows          []*Window
}

func (c *Canvas) CenterTo(index int) {

}
