package builder

import (
	"strings"
	"testing"
)

func TestHtmlElement_String(t *testing.T) {
	element := HtmlElement{name: "p", text: "hello"}

	actual := element.String()
	expected := "<p>\n\thello\n</p>\n"
	if actual != expected {
		t.Errorf("actual = %q, expected %q", actual, expected)
	}
}

func TestHtmlElement_NestedString(t *testing.T) {
	element := HtmlElement{
		name: "ul",
		elements: []HtmlElement{
			{name: "li", text: "hello"},
			{name: "li", text: "world"},
		},
	}

	actual := element.String()
	expected := `<ul>
	<li>
		hello
	</li>
	<li>
		world
	</li>
</ul>
`
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("actual = %q, expected %q", actual, expected)
	}
}

func TestHtmlBuilder_AddChildFluent(t *testing.T) {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello").AddChild("li", "world")

	actual := b.String()
	expected := `<ul>
	<li>
		hello
	</li>
	<li>
		world
	</li>
</ul>
`
	if actual != expected {
		t.Errorf("actual = %q, expected %q", actual, expected)
	}
}
