package strategy

import (
	"fmt"
	"strings"
)

type ListStrategy interface {
	Start(*strings.Builder)
	End(*strings.Builder)
	AddListItem(*strings.Builder, string)
}

type MarkdownListStrategy struct{}

func (MarkdownListStrategy) Start(*strings.Builder) {}
func (MarkdownListStrategy) End(*strings.Builder)   {}
func (MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}
func (HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}
func (HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}

// Provide static default strategies to avoid allocation
var (
	markdownStrategy = MarkdownListStrategy{}
	htmlStrategy     = HtmlListStrategy{}
)

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(strategy ListStrategy) *TextProcessor {
	return &TextProcessor{listStrategy: strategy}
}

func (t *TextProcessor) SetListStrategy(strategy ListStrategy) {
	t.listStrategy = strategy
}

func (t *TextProcessor) AppendList(items []string) {
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddListItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(markdownStrategy)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())

	tp.Reset()
	tp.SetListStrategy(htmlStrategy)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())
}
