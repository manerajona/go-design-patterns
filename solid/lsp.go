package solid

type Parallelogram interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type NaiveSquare struct {
	Rectangle
}

func NewNaiveSquare(size int) *NaiveSquare {
	sq := NaiveSquare{}
	sq.width = size
	sq.height = size
	return &sq
}

// breaks liskov substitution principle
func (s *NaiveSquare) SetWidth(width int) {
	s.width = width
	s.height = width
}

// breaks liskov substitution principle
func (s *NaiveSquare) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square struct {
	size int
}

func (s *Square) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}
