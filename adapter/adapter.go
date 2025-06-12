package adapter

import (
	"crypto/md5"
	"encoding/json"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width--
	height--
	return &VectorImage{Lines: []Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	points := owner.GetPoints()
	maxX, maxY := 0, 0
	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}
	maxX++
	maxY++

	canvas := make([][]rune, maxY)
	for y := range canvas {
		canvas[y] = make([]rune, maxX)
		for x := range canvas[y] {
			canvas[y][x] = ' '
		}
	}

	for _, point := range points {
		canvas[point.Y][point.X] = '*'
	}

	var b strings.Builder
	for _, row := range canvas {
		b.WriteString(string(row))
		b.WriteRune('\n')
	}
	return b.String()
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func rasterizeLine(line Line) []Point {
	var result []Point

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	if line.X1 == line.X2 {
		for y := top; y <= bottom; y++ {
			result = append(result, Point{X: left, Y: y})
		}
	} else if line.Y1 == line.Y2 {
		for x := left; x <= right; x++ {
			result = append(result, Point{X: x, Y: top})
		}
	}
	return result
}

func hashLine(line Line) [16]byte {
	bytes, _ := json.Marshal(line)
	return md5.Sum(bytes)
}

type vectorToRasterAdapter struct {
	points []Point
}

var pointCache = map[[16]byte][]Point{}

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	h := hashLine(line)
	if cached, ok := pointCache[h]; ok {
		a.points = append(a.points, cached...)
		return
	}

	points := rasterizeLine(line)
	pointCache[h] = points
	a.points = append(a.points, points...)
}

func (a *vectorToRasterAdapter) GetPoints() []Point {
	return a.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := &vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}
	return adapter
}
