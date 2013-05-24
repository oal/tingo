package tingo

import (
	"fmt"
	"strings"
)

type ElementInterface interface {
	setParent(parent ElementInterface) ElementInterface
	getParent() ElementInterface
	Id(i string) ElementInterface
	Class(c string) ElementInterface
	If(b bool) ElementInterface
	Before(text string) ElementInterface
	After(text string) ElementInterface
	Render() (data string)
	RenderIndent(indent string) (data string)
}

type Element struct {
	tagName    string
	attributes map[string]string
	isHidden   bool
	before     string
	children   []ElementInterface
	after      string
	parent     ElementInterface
}

func (t *Element) setParent(parent ElementInterface) ElementInterface {
	t.parent = parent
	return t
}

func (t *Element) getParent() ElementInterface {
	return t.parent
}

func (t *Element) Id(i string) ElementInterface {
	t.attributes["id"] = i
	return t
}

func (t *Element) Class(c string) ElementInterface {
	t.attributes["class"] = c
	return t
}

func (t *Element) If(b bool) ElementInterface {
	t.isHidden = !b
	return t
}

func (t *Element) Before(text string) ElementInterface {
	t.before = text
	return t
}

func (t *Element) After(text string) ElementInterface {
	t.after = text
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
			parent = parent.getParent()
			depth++
		}
		nextIndent := strings.Repeat(indent, depth)

		// Make sure text before / after is also indented
		before := t.before
		if len(before) > 0 {
			before = fmt.Sprintf("%v%v%v", indent, nextIndent, before)
			if len(children) > 0 {
				before = fmt.Sprintf("%v\n", before)
			}
		}
		after := t.after
		if len(after) > 0 {
			after = fmt.Sprintf("\n%v%v%v", indent, nextIndent, after)
			if len(children) > 0 {
				after = fmt.Sprintf("%v\n", after)
			}
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

func newElement(tagName string, children []ElementInterface) *Element {
	element := &Element{
		tagName:    tagName,
		attributes: make(map[string]string),
		children:   children,
	}
	for _, child := range children {
		child.setParent(element)
	}
	return element
}
