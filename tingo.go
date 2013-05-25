package tingo

import (
	"fmt"
	"log"
	"strings"
)

type Element struct {
	tagName    string
	parent     *Element
	attributes map[string]string
	isHidden   bool
	isVoid     bool

	children []*Element
	prepend  string
	append   string
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

// Element methods
// Available for all elements

func (el *Element) Accesskey(key string) *Element {
	el.attributes["accesskey"] = key
	return el
}

func (el *Element) Class(cls string) *Element {
	el.attributes["class"] = cls
	return el
}

func (el *Element) Id(id string) *Element {
	el.attributes["id"] = id
	return el
}

func (el *Element) Contenteditable(b bool) *Element {
	if b {
		el.attributes["contenteditable"] = "true"
	} else {
		el.attributes["contenteditable"] = "false"
	}
	return el
}

func (el *Element) Contextmenu(id string) *Element {
	el.attributes["contextmenu"] = id
	return el
}

func (el *Element) Dir(state string) *Element {
	if state == "ltr" || state == "trl" || state == "auto" {
		el.attributes["dir"] = state
	}
	return el
}

// Restricted

func (el *Element) Type(t string) *Element {
	el.attributes["type"] = t
	return el
}

func (el *Element) If(b bool) *Element {
	el.isHidden = !b
	return el
}

func (el *Element) Prepend(text string) *Element {
	if el.isVoid {
		log.Println(`WARNING: "Prepend" has no effect on void elements.`)
	}
	el.prepend = text
	return el
}

func (el *Element) Append(text string) *Element {
	if el.isVoid {
		log.Println(`WARNING: "Append" has no effect on void elements.`)
	}
	el.append = text
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
			el.prepend,
			strings.Join(children, ""),
			el.append,
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
	nextIndent := strings.Repeat(indent, depth)

	if el.isVoid {
		// Format
		return fmt.Sprintf(
			"%v<%v%v>",
			nextIndent,
			el.tagName,
			strings.Join(attributes, ""),
		)
	} else {
		children := make([]string, 0)
		for _, child := range el.children {
			children = append(children, child.RenderIndent(indent))
		}

		// Make sure text before / after is also indented
		before := el.prepend
		after := el.append

		if len(children) > 0 {
			if len(before) > 0 {
				before = fmt.Sprintf("%v%v%v\n", indent, nextIndent, before)
			}
			if len(after) > 0 {
				after = fmt.Sprintf("\n%v%v%v\n", indent, nextIndent, after)
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
		} else {
			return fmt.Sprintf(
				"%v<%v%v>%v%v</%v>",
				nextIndent,
				el.tagName,
				strings.Join(attributes, ""),
				before,
				after,
				el.tagName,
			)
		}
	}
}
