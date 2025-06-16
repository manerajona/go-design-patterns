package iterator

type Iterator[T any] interface {
	Next() (T, bool)
	HasNext() bool
	Reset()
	Current() T
}

type SliceIterator[T any] struct {
	data  []T
	index int
}

func NewSliceIterator[T any](data []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		data:  data,
		index: -1,
	}
}

func (it *SliceIterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}
	it.index++
	return it.data[it.index], true
}

func (it *SliceIterator[T]) HasNext() bool {
	return it.index+1 < len(it.data)
}

func (it *SliceIterator[T]) Reset() {
	it.index = -1
}

func (it *SliceIterator[T]) Current() T {
	if it.index < 0 || it.index >= len(it.data) {
		var zero T
		return zero
	}
	return it.data[it.index]
}
