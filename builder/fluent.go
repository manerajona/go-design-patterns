package builder

import (
	"fmt"
	"strings"
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	indentation := strings.Repeat("\t", indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", indentation, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat("\t", indent+1))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", indentation, e.name))
	return sb.String()
}

/*
BUILDER PATTERN
*/

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	htmlElement := HtmlElement{rootName, "", []HtmlElement{}}
	return &HtmlBuilder{rootName, htmlElement}
}

func (b *HtmlBuilder) AddChild(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}
