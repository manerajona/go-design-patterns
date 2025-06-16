package mediator

import (
	"testing"
)

func TestChatRoom(t *testing.T) {
	room := ChatRoom{}
	john := NewUser("John")
	jane := NewUser("Jane")
	room.Join(john)
	room.Join(jane)

	john.Announce("hi room")
	jane.Announce("oh, hey john")

	simon := NewUser("Simon")
	room.Join(simon)
	simon.Announce("hi everyone!")
	jane.DirectMessage("Simon", "glad you could join us!")

	expectedJohn := []string{
		"Room: 'Jane joins the chat'",
		"Jane: 'oh, hey john'",
		"Room: 'Simon joins the chat'",
		"Simon: 'hi everyone!'",
	}
	if len(john.chatLog) != len(expectedJohn) {
		t.Errorf("Expected %d messages in John's log, got %d", len(expectedJohn), len(john.chatLog))
	}
	for i, msg := range expectedJohn {
		if i >= len(john.chatLog) || john.chatLog[i] != msg {
			t.Errorf("John's log at %d: expected %q, got %q", i, msg, john.chatLog[i])
		}
	}

	expectedJane := []string{
		"John: 'hi room'",
		"Room: 'Simon joins the chat'",
		"Simon: 'hi everyone!'",
	}
	if len(jane.chatLog) != len(expectedJane) {
		t.Errorf("Expected %d messages in Jane's log, got %d", len(expectedJane), len(jane.chatLog))
	}
	for i, msg := range expectedJane {
		if i >= len(jane.chatLog) || jane.chatLog[i] != msg {
			t.Errorf("Jane's log at %d: expected %q, got %q", i, msg, jane.chatLog[i])
		}
	}

	expectedSimon := []string{
		"Jane: 'glad you could join us!'",
	}
	if len(simon.chatLog) != len(expectedSimon) {
		t.Errorf("Expected %d messages in Simon's log, got %d", len(expectedSimon), len(simon.chatLog))
	}
	for i, msg := range expectedSimon {
		if i >= len(simon.chatLog) || simon.chatLog[i] != msg {
			t.Errorf("Simon's log at %d: expected %q, got %q", i, msg, simon.chatLog[i])
		}
	}
}
