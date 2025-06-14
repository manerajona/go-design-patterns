package proxy

import (
	"testing"
)

func TestLazyBitmap(t *testing.T) {
	bmp := NewLazyBitmap("demo.png")

	if bmp.bitmap != nil {
		t.Errorf("Bitmap already loaded")
	}

	bmp.Draw()
	if bmp.bitmap == nil {
		t.Errorf("Bitmap not loaded")
	}
}
