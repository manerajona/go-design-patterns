package interpreter

import (
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{"1+2", []Token{{Int, "1"}, {Plus, "+"}, {Int, "2"}}},
		{"34-5", []Token{{Int, "34"}, {Minus, "-"}, {Int, "5"}}},
		{"(3+4)", []Token{{Lparen, "("}, {Int, "3"}, {Plus, "+"}, {Int, "4"}, {Rparen, ")"}}},
		{" 18 + (21-9) ", []Token{
			{Int, "18"}, {Plus, "+"}, {Lparen, "("}, {Int, "21"}, {Minus, "-"}, {Int, "9"}, {Rparen, ")"},
		}},
	}

	for _, tt := range tests {
		got := Lex(tt.input)
		if len(got) != len(tt.expected) {
			t.Errorf("Lex(%q): expected %v tokens, got %v", tt.input, len(tt.expected), len(got))
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("Lex(%q): at %d expected %v, got %v", tt.input, i, tt.expected[i], got[i])
			}
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1+2", 3},
		{"5-2", 3},
		{"(3+4)", 7},
		{"(20-7)", 13},
		{"(13+4)-(12+1)", 4},
		{"18+(21-9)", 30},
	}

	for _, tt := range tests {
		tokens := Lex(tt.input)
		root := Parse(tokens)
		got := root.Value()
		if got != tt.expected {
			t.Errorf("Parse(%q): got %d, want %d", tt.input, got, tt.expected)
		}
	}
}
