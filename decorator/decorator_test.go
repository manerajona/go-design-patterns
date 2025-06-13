package decorator

import (
	"strings"
	"testing"
)

func TestCircle_RenderAndResize(t *testing.T) {
	c := &Circle{2}
	want := "Circle of radius 2"
	if got := c.Render(); got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
	c.Resize(2)
	want = "Circle of radius 4"
	if got := c.Render(); got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestSquare_Render(t *testing.T) {
	s := &Square{3}
	want := "Square with side 3"
	if got := s.Render(); got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestColoredCircle_Render(t *testing.T) {
	circle := &Circle{2}
	redCircle := &ColoredShape{Shape: circle, Color: "Red"}
	got := redCircle.Render()
	if !strings.Contains(got, "Circle of radius 2.000000") || !strings.Contains(got, "Red") {
		t.Errorf("Expected description with radius and color, got %q", got)
	}
}

func TestTransparentCircle_Render(t *testing.T) {
	circle := &Circle{2}
	redCircle := &ColoredShape{Shape: circle, Color: "Red"}
	rhsCircle := &TransparentShape{Shape: redCircle, Transparency: 0.5}
	got := rhsCircle.Render()
	if !strings.Contains(got, "Circle of radius 2.000000") ||
		!strings.Contains(got, "Red") ||
		!strings.Contains(got, "0.500000") {
		t.Errorf("Expected description with radius, color, and transparency, got %q", got)
	}
}

func TestColoredSquare_Render(t *testing.T) {
	square := &Square{2}
	redSquare := &ColoredShape{Shape: square, Color: "Red"}
	got := redSquare.Render()
	if !strings.Contains(got, "Square with side 2.000000") || !strings.Contains(got, "Red") {
		t.Errorf("Expected description with side and color, got %q", got)
	}
}

func TestTransparentSquare_Render(t *testing.T) {
	square := &Square{2}
	redSquare := &ColoredShape{Shape: square, Color: "Red"}
	rhsSquare := &TransparentShape{Shape: redSquare, Transparency: 0.5}
	got := rhsSquare.Render()
	if !strings.Contains(got, "Square with side 2.000000") ||
		!strings.Contains(got, "Red") ||
		!strings.Contains(got, "0.500000") {
		t.Errorf("Expected description with side, color, and transparency, got %q", got)
	}
}
