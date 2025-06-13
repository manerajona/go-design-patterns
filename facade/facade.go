package facade

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: make([]rune, width*height),
	}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(index)
}

// Console is a facade over buffers and viewports
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	return &Console{
		buffers:   []*Buffer{b},
		viewports: []*Viewport{v},
		offset:    0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[c.offset].GetCharacterAt(index)
}
