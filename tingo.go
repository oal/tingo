package tingo

import (
	"fmt"
	"html"
	"strings"
)

type Element struct {
	tagName    string
	parent     *Element
	attributes map[string]string
	isHidden   bool
	isVoid     bool
	isSafe     bool

	children    []*Element
	textPrepend string
	textAppend  string
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

func (el *Element) Draggable(b bool) *Element {
	if b {
		el.attributes["draggable"] = "true"
	} else {
		el.attributes["draggable"] = "false"
	}
	return el
}

func (el *Element) Dropzone(state string) *Element {
	if state == "copy" || state == "move" || state == "link" {
		el.attributes["dropzone"] = state
	}
	return el
}

func (el *Element) Hidden(b bool) *Element {
	if b {
		el.attributes["draggable"] = "true"
	} else {
		el.attributes["draggable"] = "false"
	}
	return el
}

func (el *Element) Id(id string) *Element {
	el.attributes["id"] = id
	return el
}

func (el *Element) Lang(lang string) *Element {
	el.attributes["lang"] = lang
	return el
}

func (el *Element) Spellcheck(b bool) *Element {
	if b {
		el.attributes["spellcheck"] = "true"
	} else {
		el.attributes["spellcheck"] = "false"
	}
	return el
}

func (el *Element) Style(style string) *Element {
	el.attributes["style"] = style
	return el
}

func (el *Element) Tabindex(index int) *Element {
	el.attributes["tabindex"] = fmt.Sprintf("%v", index)
	return el
}

func (el *Element) Title(title string) *Element {
	el.attributes["title"] = title
	return el
}

func (el *Element) Translate(b bool) *Element {
	if b {
		el.attributes["translate"] = "yes"
	} else {
		el.attributes["translate"] = "no"
	}
	return el
}

// Restricted

func (el *Element) Type(t string) *Element {
	el.attributes["type"] = t
	return el
}

// Logic and additional methods

func (el *Element) Safe(b bool) *Element {
	// All text and attributes will be escaped by default. Call Safe(true) if you trust the input.
	el.isSafe = b

	var walk func(*Element)
	walk = func(element *Element) {
		for _, child := range element.children {
			child.isSafe = b
			walk(child)
		}
	}
	walk(el)
	return el
}

func (el *Element) If(b bool) *Element {
	// If b is false, this element won't be rendered.
	el.isHidden = !b
	return el
}

func (el *Element) TextPrepend(text string) *Element {
	// Adds text before this element's children. Has no effect on void elements like <br>.
	el.textPrepend = text
	return el
}

func (el *Element) TextAppend(text string) *Element {
	// Adds text after this element's children. Has no effect on void elements like <br>.
	el.textAppend = text
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
			el.textPrepend,
			strings.Join(children, ""),
			el.textAppend,
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

		textPrepend := el.textPrepend
		textAppend := el.textAppend

		if !el.isSafe {
			textPrepend = html.EscapeString(textPrepend)
			textAppend = html.EscapeString(textAppend)
		}

		// If element has children, indent text and children.
		// Otherwise, add appended/prepended text on the same line like <a>text</a>.
		if len(children) > 0 {
			// Make sure prepended / appended text is also indented
			if len(textPrepend) > 0 {
				textPrepend = fmt.Sprintf("%v%v%v\n", indent, nextIndent, textPrepend)
			}
			if len(textAppend) > 0 {
				textAppend = fmt.Sprintf("\n%v%v%v\n", indent, nextIndent, textAppend)
			}
			return fmt.Sprintf(
				"%v<%v%v>\n%v%v%v\n%v</%v>",
				nextIndent,
				el.tagName,
				strings.Join(attributes, ""),
				textPrepend,
				strings.Join(children, "\n"),
				textAppend,
				nextIndent,
				el.tagName,
			)
		} else {
			return fmt.Sprintf(
				"%v<%v%v>%v%v</%v>",
				nextIndent,
				el.tagName,
				strings.Join(attributes, ""),
				textPrepend,
				textAppend,
				el.tagName,
			)
		}
	}
}
