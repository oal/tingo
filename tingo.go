package tingo

import (
	"fmt"
	"strings"
)

type Element struct {
	tagName    string
	attributes map[string]string
	isHidden   bool
	before     string
	children   []*Element
	after      string
	parent     *Element
}

type ElementFunc func(elements ...*Element) *Element

func (t *Element) Id(i string) *Element {
	t.attributes["id"] = i
	return t
}

func (t *Element) Class(c string) *Element {
	t.attributes["class"] = c
	return t
}

func (t *Element) If(b bool) *Element {
	t.isHidden = !b
	return t
}

func (t *Element) Before(text string) *Element {
	t.before = text
	return t
}

func (t *Element) After(text string) *Element {
	t.after = text
	return t
}

func (t *Element) Loop(f ElementFunc, values []string) *Element {
	for _, value := range values {
		child := f().Before(fmt.Sprintf("%v", value))
		child.parent = t
		t.children = append(t.children, child)
	}
	return t
}

func (t *Element) Render() (data string) {
	if !t.isHidden {
		attributes := make([]string, 0)
		for key, val := range t.attributes {
			attributes = append(attributes, fmt.Sprintf(` %v="%v"`, key, val))
		}

		children := make([]string, 0)
		for _, child := range t.children {
			children = append(children, child.Render())
		}

		// Format
		data = fmt.Sprintf(
			`<%v%v>%v%v%v</%v>`,
			t.tagName,
			strings.Join(attributes, ""),
			t.before,
			strings.Join(children, ""),
			t.after,
			t.tagName,
		)
	}
	return
}

func (t *Element) RenderIndent(indent string) (data string) {
	if !t.isHidden {
		attributes := make([]string, 0)
		for key, val := range t.attributes {
			attributes = append(attributes, fmt.Sprintf(` %v="%v"`, key, val))
		}

		children := make([]string, 0)
		for _, child := range t.children {
			children = append(children, child.RenderIndent(indent))
		}

		// Calculate depth to find indention level
		depth := 0
		parent := t.parent
		for parent != nil {
			parent = parent.parent
			depth++
		}
		nextIndent := strings.Repeat(indent, depth)

		// Make sure text before / after is also indented
		before := t.before
		if len(before) > 0 {
			before = fmt.Sprintf("%v%v%v\n", indent, nextIndent, before)
		}
		after := t.after
		if len(after) > 0 {
			after = fmt.Sprintf("\n%v%v%v", indent, nextIndent, after)
		}

		// Format with indentation
		data = fmt.Sprintf(
			"%v<%v%v>\n%v%v%v\n%v</%v>",
			nextIndent,
			t.tagName,
			strings.Join(attributes, ""),
			before,
			strings.Join(children, "\n"),
			after,
			nextIndent,
			t.tagName,
		)
	}
	return
}
