package observer

import (
	"container/list"
)

type Observer[T any] interface {
	Notify(T)
}

type Observable[T any] interface {
	Subscribe(Observer[T])
	Unsubscribe(Observer[T])
	Broadcast(T)
}

type observable[T any] struct {
	subscribers *list.List
}

func NewObservable[T any]() Observable[T] {
	return &observable[T]{subscribers: list.New()}
}

func (o *observable[T]) Subscribe(observer Observer[T]) {
	o.subscribers.PushBack(observer)
}

func (o *observable[T]) Unsubscribe(observer Observer[T]) {
	for it := o.subscribers.Front(); it != nil; it = it.Next() {
		if it.Value.(Observer[T]) == observer {
			o.subscribers.Remove(it)
			break
		}
	}
}

func (o *observable[T]) Broadcast(event T) {
	for it := o.subscribers.Front(); it != nil; it = it.Next() {
		it.Value.(Observer[T]).Notify(event)
	}
}

type Event struct {
	Message string
}

type EventObserver struct {
	name     string
	messages []string
}

func NewEventObserver(name string) *EventObserver {
	return &EventObserver{name: name}
}

func (c *EventObserver) Notify(event Event) {
	c.messages = append(c.messages, event.Message)
}
