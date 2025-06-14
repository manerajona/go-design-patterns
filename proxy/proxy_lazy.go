package proxy

import "fmt"

type Image interface {
	Draw()
}

type bitmap struct {
	filename string
}

func (b *bitmap) Draw() {
	fmt.Println("Drawing", b.filename)
}

type LazyBitmap struct {
	filename string
	*bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		fmt.Println("Loading", l.filename)
		l.bitmap = &bitmap{filename: l.filename}
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}
