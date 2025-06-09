package solid

import (
	"testing"
)

func testBehaviour(t *testing.T, p Parallelogram, newHeight int) {
	expectedArea := newHeight * p.GetWidth()
	p.SetHeight(newHeight)
	actualArea := p.GetWidth() * p.GetHeight()
	if actualArea != expectedArea {
		t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
	}
}

func TestRectangleBehaviour_Pass(t *testing.T) {
	rectangle := &Rectangle{2, 3}
	testBehaviour(t, rectangle, 10)
}

func TestNaiveSquareBehaviour_Fail(t *testing.T) {
	square := NewNaiveSquare(5)
	testBehaviour(t, square, 10)
}

func TestSquareBehaviour_Pass(t *testing.T) {
	square := &Square{size: 5}
	testBehaviour(t, square.Rectangle(), 10)
}
