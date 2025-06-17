package strategy

import (
	"strings"
	"testing"
)

func TestMarkdownListStrategy(t *testing.T) {
	builder := &strings.Builder{}
	strategy := MarkdownListStrategy{}

	strategy.Start(builder)
	strategy.AddListItem(builder, "item1")
	strategy.AddListItem(builder, "item2")
	strategy.End(builder)

	expected := " * item1\n * item2\n"
	if builder.String() != expected {
		t.Errorf("MarkdownListStrategy output incorrect.\nExpected:\n%s\nGot:\n%s", expected, builder.String())
	}
}

func TestHtmlListStrategy(t *testing.T) {
	builder := &strings.Builder{}
	strategy := HtmlListStrategy{}

	strategy.Start(builder)
	strategy.AddListItem(builder, "item1")
	strategy.AddListItem(builder, "item2")
	strategy.End(builder)

	expected := "<ul>\n  <li>item1</li>\n  <li>item2</li>\n</ul>\n"
	if builder.String() != expected {
		t.Errorf("HtmlListStrategy output incorrect.\nExpected:\n%s\nGot:\n%s", expected, builder.String())
	}
}

func TestTextProcessor_Markdown(t *testing.T) {
	tp := NewTextProcessor(markdownStrategy)
	items := []string{"foo", "bar"}
	tp.AppendList(items)

	expected := " * foo\n * bar\n"
	if tp.String() != expected {
		t.Errorf("TextProcessor markdown output incorrect.\nExpected:\n%s\nGot:\n%s", expected, tp.String())
	}
}

func TestTextProcessor_Html(t *testing.T) {
	tp := NewTextProcessor(htmlStrategy)
	items := []string{"foo", "bar"}
	tp.AppendList(items)

	expected := "<ul>\n  <li>foo</li>\n  <li>bar</li>\n</ul>\n"
	if tp.String() != expected {
		t.Errorf("TextProcessor html output incorrect.\nExpected:\n%s\nGot:\n%s", expected, tp.String())
	}
}

func TestTextProcessor_Reset(t *testing.T) {
	tp := NewTextProcessor(markdownStrategy)
	tp.AppendList([]string{"reset"})
	tp.Reset()
	if tp.String() != "" {
		t.Errorf("TextProcessor Reset failed, expected empty string, got: %s", tp.String())
	}
}

func TestTextProcessor_SetListStrategy(t *testing.T) {
	tp := NewTextProcessor(markdownStrategy)
	tp.AppendList([]string{"foo"})
	tp.Reset()
	tp.SetListStrategy(htmlStrategy)
	tp.AppendList([]string{"foo"})
	expected := "<ul>\n  <li>foo</li>\n</ul>\n"
	if tp.String() != expected {
		t.Errorf("TextProcessor SetListStrategy output incorrect.\nExpected:\n%s\nGot:\n%s", expected, tp.String())
	}
}
