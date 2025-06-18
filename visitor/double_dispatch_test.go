package visitor

import (
	"testing"
)

func TestPrinterExpressionVisitor_AdditionExpression(t *testing.T) {
	// Expression: (1+(2+3))
	e := &AdditionExpressionV2{
		left: &DoubleExpressionV2{1},
		right: &AdditionExpressionV2{
			left:  &DoubleExpressionV2{2},
			right: &DoubleExpressionV2{3},
		},
	}
	visitor := NewPrinterExpressionVisitor()
	e.Accept(visitor)
	got := visitor.(*PrinterExpressionVisitor).String()
	want := "(1+(2+3))"
	if got != want {
		t.Errorf("expected '%s', got '%s'", want, got)
	}
}

func TestPrinterExpressionVisitor_DoubleExpression(t *testing.T) {
	e := &DoubleExpressionV2{42.5}
	visitor := NewPrinterExpressionVisitor()
	e.Accept(visitor)
	got := visitor.(*PrinterExpressionVisitor).String()
	want := "42.5"
	if got != want {
		t.Errorf("expected '%s', got '%s'", want, got)
	}
}
