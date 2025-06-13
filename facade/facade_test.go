package facade

import "testing"

func TestBuffer_At(t *testing.T) {
	b := NewBuffer(5, 2)
	b.buffer[3] = 'X'
	if ch := b.At(3); ch != 'X' {
		t.Errorf("Expected 'X', got %c", ch)
	}
}

func TestViewport_GetCharacterAt(t *testing.T) {
	b := NewBuffer(4, 1)
	b.buffer[2] = 'A'
	v := NewViewport(b)
	if ch := v.GetCharacterAt(2); ch != 'A' {
		t.Errorf("Expected 'A', got %c", ch)
	}
}

func TestConsole_GetCharacterAt(t *testing.T) {
	c := NewConsole()
	c.buffers[0].buffer[5] = 'Z'
	ch := c.GetCharacterAt(5)
	if ch != 'Z' {
		t.Errorf("Expected 'Z', got %c", ch)
	}
}

func TestConsole_MultipleViewportsAndBuffers(t *testing.T) {
	console := &Console{}

	// Add 2 buffers for each of 2 viewports
	buffer1 := NewBuffer(3, 1)
	buffer2 := NewBuffer(3, 1)
	buffer3 := NewBuffer(3, 1)
	buffer4 := NewBuffer(3, 1)

	// Fill buffers with unique runes
	buffer1.buffer[0] = 'A'
	buffer1.buffer[1] = 'B'
	buffer1.buffer[2] = 'C'

	buffer2.buffer[0] = 'D'
	buffer2.buffer[1] = 'E'
	buffer2.buffer[2] = 'F'

	buffer3.buffer[0] = 'G'
	buffer3.buffer[1] = 'H'
	buffer3.buffer[2] = 'I'

	buffer4.buffer[0] = 'J'
	buffer4.buffer[1] = 'K'
	buffer4.buffer[2] = 'L'

	console.buffers = []*Buffer{buffer1, buffer2, buffer3, buffer4}

	// Each viewport uses one buffer
	viewport1 := NewViewport(buffer1)
	viewport2 := NewViewport(buffer2)
	viewport3 := NewViewport(buffer3)
	viewport4 := NewViewport(buffer4)
	console.viewports = []*Viewport{viewport1, viewport2, viewport3, viewport4}

	// Assert each viewport returns correct rune at specific position
	if ch := console.GetCharacterAt(0); ch != 'A' {
		t.Errorf("Expected 'A' in viewport 0, got %c", ch)
	}
	if ch := console.GetCharacterAt(2); ch != 'C' {
		t.Errorf("Expected 'C' in viewport 0, got %c", ch)
	}
	console.offset = 1
	if ch := console.GetCharacterAt(2); ch != 'F' {
		t.Errorf("Expected 'F' in viewport 1, got %c", ch)
	}
	console.offset = 2
	if ch := console.GetCharacterAt(2); ch != 'I' {
		t.Errorf("Expected 'I' in viewport 2, got %c", ch)
	}
	console.offset = 3
	if ch := console.GetCharacterAt(2); ch != 'L' {
		t.Errorf("Expected 'L' in viewport 3, got %c", ch)
	}
}
