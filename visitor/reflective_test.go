package visitor

import (
	"strings"
	"testing"
)

func TestPrint_AdditionExpression(t *testing.T) {
	// (1+(2+3))
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := strings.Builder{}
	Print(e, &sb)
	got := sb.String()
	want := "(1+(2+3))"
	if got != want {
		t.Errorf("expected '%s', got '%s'", want, got)
	}
}

func TestPrint_DoubleExpression(t *testing.T) {
	e := &DoubleExpression{42.5}
	sb := strings.Builder{}
	Print(e, &sb)
	got := sb.String()
	want := "42.5"
	if got != want {
		t.Errorf("expected '%s', got '%s'", want, got)
	}
}
