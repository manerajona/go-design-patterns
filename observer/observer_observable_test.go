package observer

import (
	"strings"
	"testing"
)

func TestObservable_SubscribeAndBroadcast(t *testing.T) {
	obs := NewObservable[Event]()

	a := NewEventObserver("A")
	b := NewEventObserver("B")

	obs.Subscribe(a)
	obs.Subscribe(b)

	// Broadcast an event
	testMsg := "Test Message"
	obs.Broadcast(Event{Message: testMsg})

	if len(a.messages) != 1 || a.messages[0] != testMsg {
		t.Errorf("Consumer A did not receive correct message, got: %v", a.messages)
	}
	if len(b.messages) != 1 || b.messages[0] != testMsg {
		t.Errorf("Consumer B did not receive correct message, got: %v", b.messages)
	}
}

func TestObservable_Unsubscribe(t *testing.T) {
	obs := NewObservable[Event]()

	a := NewEventObserver("A")
	b := NewEventObserver("B")

	obs.Subscribe(a)
	obs.Subscribe(b)

	obs.Broadcast(Event{Message: "First"})
	obs.Unsubscribe(a)
	obs.Broadcast(Event{Message: "Second"})

	// A should only have "First", B should have both
	expectedA := []string{"First"}
	expectedB := []string{"First", "Second"}

	if strings.Join(a.messages, ",") != strings.Join(expectedA, ",") {
		t.Errorf("Consumer A did not receive expected messages after unsubscribe, got: %v", a.messages)
	}
	if strings.Join(b.messages, ",") != strings.Join(expectedB, ",") {
		t.Errorf("Consumer B did not receive expected messages after unsubscribe, got: %v", b.messages)
	}
}
