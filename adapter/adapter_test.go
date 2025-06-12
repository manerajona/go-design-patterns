package adapter

import (
	"fmt"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	vi := NewRectangle(4, 3)
	if len(vi.Lines) != 4 {
		t.Fatalf("Expected 4 lines for rectangle, got %d", len(vi.Lines))
	}
}

func TestVectorToRaster_DrawRectangle(t *testing.T) {
	vi := NewRectangle(6, 4)
	raster := VectorToRaster(vi)
	output := DrawPoints(raster)
	fmt.Println(output)
	expected :=
		`******` + "\n" +
			`*    *` + "\n" +
			`*    *` + "\n" +
			`******` + "\n"
	if output != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, output)
	}
}

func TestVectorToRaster_DrawRectangleCached(t *testing.T) {
	vi := NewRectangle(3, 2)
	_ = VectorToRaster(vi)
	countBefore := len(pointCache)
	_ = VectorToRaster(vi)
	countAfter := len(pointCache)
	if countBefore != countAfter {
		t.Errorf("Point cache should not grow after repeated rasterizations. Before: %d, After: %d", countBefore, countAfter)
	}
}

func TestDrawPoints_DrawPoints(t *testing.T) {
	points := []Point{{0, 0}, {1, 1}, {2, 2}}
	raster := &vectorToRasterAdapter{points: points}
	output := DrawPoints(raster)
	fmt.Println(output)
	expected :=
		`*  ` + "\n" +
			` * ` + "\n" +
			`  *` + "\n"
	if output != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, output)
	}
}
