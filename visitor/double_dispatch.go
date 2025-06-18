package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(*DoubleExpressionV2)
	VisitAdditionExpression(*AdditionExpressionV2)
}

type ExpressionV2 interface {
	Accept(ExpressionVisitor)
}

type DoubleExpressionV2 struct {
	value float64
}

func (d *DoubleExpressionV2) Accept(visitor ExpressionVisitor) {
	visitor.VisitDoubleExpression(d)
}

type AdditionExpressionV2 struct {
	left, right ExpressionV2
}

func (a *AdditionExpressionV2) Accept(visitor ExpressionVisitor) {
	visitor.VisitAdditionExpression(a)
}

type PrinterExpressionVisitor struct {
	sb strings.Builder
}

func NewPrinterExpressionVisitor() ExpressionVisitor {
	return &PrinterExpressionVisitor{strings.Builder{}}
}

func (visitor *PrinterExpressionVisitor) VisitDoubleExpression(expression *DoubleExpressionV2) {
	visitor.sb.WriteString(fmt.Sprintf("%g", expression.value))
}

func (visitor *PrinterExpressionVisitor) VisitAdditionExpression(expression *AdditionExpressionV2) {
	visitor.sb.WriteString("(")
	expression.left.Accept(visitor)
	visitor.sb.WriteString("+")
	expression.right.Accept(visitor)
	visitor.sb.WriteString(")")
}

func (visitor *PrinterExpressionVisitor) String() string {
	return visitor.sb.String()
}
