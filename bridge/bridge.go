package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	Called bool
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	v.Called = true
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}
