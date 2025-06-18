package visitor

import (
	"fmt"
	"strings"
)

type Expression interface {
	// Marker interface
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(expression Expression, visitor *strings.Builder) {
	if e, ok := expression.(*DoubleExpression); ok {
		visitor.WriteString(fmt.Sprintf("%g", e.value))
	} else if e, ok := expression.(*AdditionExpression); ok {
		visitor.WriteString("(")
		Print(e.left, visitor)
		visitor.WriteString("+")
		Print(e.right, visitor)
		visitor.WriteString(")")
	}
	// breaks OCP
	// will work incorrectly on missing case
}
