package decorator

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// ColoredShape extends Shape functionality
type ColoredShape struct {
	Shape Shape // embedded
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has color %s", c.Shape.Render(), c.Color)
}

// TransparentShape extends Shape functionality
type TransparentShape struct {
	Shape        Shape // embedded
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f transparency", t.Shape.Render(), t.Transparency)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
