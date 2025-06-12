package bridge

import "testing"

func TestCircle_DrawWithVectorRenderer(t *testing.T) {
	tr := &VectorRenderer{}
	circle := NewCircle(tr, 7.5)
	circle.Draw()
	if !tr.Called {
		t.Errorf("Expected RenderCircle to be called with radius 7.5, got %+v", tr)
	}
}

func TestCircle_DrawWithRasterRenderer(t *testing.T) {
	tr := &VectorRenderer{}
	circle := NewCircle(tr, 3.2)
	circle.Draw()
	if !tr.Called {
		t.Errorf("Expected RenderCircle to be called with radius 3.2, got %+v", tr)
	}
}
