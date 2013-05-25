package tingo

import (
	"fmt"
	"strings"
)

type Element struct {
	tagName    string
	parent     *Element
	attributes map[string]string
	isHidden   bool
	isVoid     bool

	children []*Element
	before   string
	after    string
}

func newVoidElement(tagName string) *Element {
	element := &Element{
		tagName:    tagName,
		attributes: make(map[string]string),
		isVoid:     true,
	}
	return element
}

func newElement(tagName string, children []*Element) *Element {
	element := &Element{
		tagName:    tagName,
		attributes: make(map[string]string),
		children:   children,
	}
	for _, child := range children {
		child.parent = element
	}
	return element
}

func (el *Element) Id(i string) *Element {
	el.attributes["id"] = i
	return el
}

func (el *Element) Class(c string) *Element {
	el.attributes["class"] = c
	return el
}

func (el *Element) Type(t string) *Element {
	el.attributes["type"] = t
	return el
}

func (el *Element) If(b bool) *Element {
	el.isHidden = !b
	return el
}

func (el *Element) Before(text string) *Element {
	el.before = text
	return el
}

func (el *Element) After(text string) *Element {
	el.after = text
	return el
}

func (el *Element) Render() string {
	if el.isHidden {
		return ""
	}

	attributes := make([]string, 0)
	for key, val := range el.attributes {
		attributes = append(attributes, fmt.Sprintf(` %v="%v"`, key, val))
	}

	if el.isVoid {
		// Format
		return fmt.Sprintf(
			`<%v%v>`,
			el.tagName,
			strings.Join(attributes, ""),
		)
	} else {
		children := make([]string, 0)
		for _, child := range el.children {
			children = append(children, child.Render())
		}

		return fmt.Sprintf(
			`<%v%v>%v%v%v</%v>`,
			el.tagName,
			strings.Join(attributes, ""),
			el.before,
			strings.Join(children, ""),
			el.after,
			el.tagName,
		)
	}
}

func (el *Element) RenderIndent(indent string) string {
	if el.isHidden {
		return ""
	}

	attributes := make([]string, 0)
	for key, val := range el.attributes {
		attributes = append(attributes, fmt.Sprintf(` %v="%v"`, key, val))
	}

	// Calculate depth to find indention level
	depth := 0
	parent := el.parent
	for parent != nil {
		parent = parent.parent
		depth++
	}
	fmt.Println(depth)
	nextIndent := strings.Repeat(indent, depth)

	if el.isVoid {
		// Format
		return fmt.Sprintf(
			"%v<%v%v>",
			indent,
			el.tagName,
			strings.Join(attributes, ""),
		)
	} else {
		children := make([]string, 0)
		for _, child := range el.children {
			children = append(children, child.Render())
		}

		// Make sure text before / after is also indented
		before := el.before
		if len(before) > 0 {
			before = fmt.Sprintf("%v%v%v", indent, nextIndent, before)
			if len(children) > 0 {
				before = fmt.Sprintf("%v\n", before)
			}
		}
		after := el.after
		if len(after) > 0 {
			after = fmt.Sprintf("\n%v%v%v", indent, nextIndent, after)
			if len(children) > 0 {
				after = fmt.Sprintf("%v\n", after)
			}
		}

		return fmt.Sprintf(
			"%v<%v%v>\n%v%v%v\n%v</%v>",
			nextIndent,
			el.tagName,
			strings.Join(attributes, ""),
			before,
			strings.Join(children, "\n"),
			after,
			nextIndent,
			el.tagName,
		)
	}
}
