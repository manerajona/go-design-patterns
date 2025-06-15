package interpreter

import (
	"strconv"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

type Operation int

const (
	Addition Operation = iota
	Subtraction
)

type BinaryOperation struct {
	Op    Operation
	Left  Element
	Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Op {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()
	}
	panic("unsupported operation")
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type TokenType
	Text string
}

func Lex(input string) []Token {
	var tokens []Token
	i := 0
	for i < len(input) {
		c := input[i]
		switch {
		case unicode.IsSpace(rune(c)):
			i++
		case c == '+':
			tokens = append(tokens, Token{Plus, "+"})
			i++
		case c == '-':
			tokens = append(tokens, Token{Minus, "-"})
			i++
		case c == '(':
			tokens = append(tokens, Token{Lparen, "("})
			i++
		case c == ')':
			tokens = append(tokens, Token{Rparen, ")"})
			i++
		case unicode.IsDigit(rune(c)):
			j := i
			for j < len(input) && unicode.IsDigit(rune(input[j])) {
				j++
			}
			tokens = append(tokens, Token{Int, input[i:j]})
			i = j
		default:
			i++
		}
	}
	return tokens
}

func Parse(tokens []Token) Element {
	var lhs Element
	var op Operation
	var rhs Element
	i := 0
	for i < len(tokens) {
		tok := tokens[i]
		switch tok.Type {
		case Int:
			val, _ := strconv.Atoi(tok.Text)
			if lhs == nil {
				lhs = &Integer{val}
			} else {
				rhs = &Integer{val}
			}
			i++
		case Plus:
			op = Addition
			i++
		case Minus:
			op = Subtraction
			i++
		case Lparen:
			j := i + 1
			balance := 1
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Lparen {
					balance++
				} else if tokens[j].Type == Rparen {
					balance--
					if balance == 0 {
						break
					}
				}
			}
			subexpr := Parse(tokens[i+1 : j])
			if lhs == nil {
				lhs = subexpr
			} else {
				rhs = subexpr
			}
			i = j + 1
		default:
			i++
		}
	}
	if rhs == nil {
		return lhs
	}
	return &BinaryOperation{Op: op, Left: lhs, Right: rhs}
}
