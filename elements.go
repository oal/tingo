package tingo

func Colgroup(elements ...*Element) *Element {
	// Table column group

	element := &Element{
		tagName:    "colgroup",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Figure(elements ...*Element) *Element {
	// Figure with optional caption (HTML5)

	element := &Element{
		tagName:    "figure",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Dd(elements ...*Element) *Element {
	// Description or value

	element := &Element{
		tagName:    "dd",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Rp(elements ...*Element) *Element {
	// Ruby parenthesis (HTML5)

	element := &Element{
		tagName:    "rp",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Source(elements ...*Element) *Element {
	// Media source (HTML5)

	element := &Element{
		tagName:    "source",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func U(elements ...*Element) *Element {
	// Offset text conventionally styled with an underline

	element := &Element{
		tagName:    "u",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Meta(elements ...*Element) *Element {
	// Metadata

	element := &Element{
		tagName:    "meta",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Ruby(elements ...*Element) *Element {
	// Ruby annotation (HTML5)

	element := &Element{
		tagName:    "ruby",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Area(elements ...*Element) *Element {
	// Image-map hyperlink

	element := &Element{
		tagName:    "area",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Aside(elements ...*Element) *Element {
	// Tangential content (HTML5)

	element := &Element{
		tagName:    "aside",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Img(elements ...*Element) *Element {
	// Image

	element := &Element{
		tagName:    "img",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Keygen(elements ...*Element) *Element {
	// Key-pair generator/input control (HTML5)

	element := &Element{
		tagName:    "keygen",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Base(elements ...*Element) *Element {
	// Base URL

	element := &Element{
		tagName:    "base",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Em(elements ...*Element) *Element {
	// Emphatic stress

	element := &Element{
		tagName:    "em",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Wbr(elements ...*Element) *Element {
	// Line-break opportunity (HTML5)

	element := &Element{
		tagName:    "wbr",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Article(elements ...*Element) *Element {
	// Article (HTML5)

	element := &Element{
		tagName:    "article",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Del(elements ...*Element) *Element {
	// Deleted text

	element := &Element{
		tagName:    "del",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Iframe(elements ...*Element) *Element {
	// Nested browsing context (inline frame)

	element := &Element{
		tagName:    "iframe",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Strong(elements ...*Element) *Element {
	// Strong importance

	element := &Element{
		tagName:    "strong",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Address(elements ...*Element) *Element {
	// Contact information

	element := &Element{
		tagName:    "address",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Code(elements ...*Element) *Element {
	// Code fragment

	element := &Element{
		tagName:    "code",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Form(elements ...*Element) *Element {
	// User-submittable form

	element := &Element{
		tagName:    "form",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Object(elements ...*Element) *Element {
	// Generic external content

	element := &Element{
		tagName:    "object",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Rt(elements ...*Element) *Element {
	// Ruby text (HTML5)

	element := &Element{
		tagName:    "rt",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Col(elements ...*Element) *Element {
	// Table column

	element := &Element{
		tagName:    "col",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Command(elements ...*Element) *Element {
	// Command (HTML5)

	element := &Element{
		tagName:    "command",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Hr(elements ...*Element) *Element {
	// Thematic break

	element := &Element{
		tagName:    "hr",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Button(elements ...*Element) *Element {
	// Button

	element := &Element{
		tagName:    "button",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Figcaption(elements ...*Element) *Element {
	// Figure caption (HTML5)

	element := &Element{
		tagName:    "figcaption",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Input(elements ...*Element) *Element {
	// Input control

	element := &Element{
		tagName:    "input",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Sup(elements ...*Element) *Element {
	// Superscript

	element := &Element{
		tagName:    "sup",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Blockquote(elements ...*Element) *Element {
	// Block quotation

	element := &Element{
		tagName:    "blockquote",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func I(elements ...*Element) *Element {
	// Offset text conventionally styled in italic

	element := &Element{
		tagName:    "i",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Legend(elements ...*Element) *Element {
	// Title or explanatory caption

	element := &Element{
		tagName:    "legend",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Td(elements ...*Element) *Element {
	// Table cell

	element := &Element{
		tagName:    "td",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Caption(elements ...*Element) *Element {
	// Table title

	element := &Element{
		tagName:    "caption",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Textarea(elements ...*Element) *Element {
	// Text input area

	element := &Element{
		tagName:    "textarea",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Div(elements ...*Element) *Element {
	// Generic flow container

	element := &Element{
		tagName:    "div",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H5(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h5",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Meter(elements ...*Element) *Element {
	// Scalar gauge (HTML5)

	element := &Element{
		tagName:    "meter",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Section(elements ...*Element) *Element {
	// Section (HTML5)

	element := &Element{
		tagName:    "section",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Details(elements ...*Element) *Element {
	// Control for additional on-demand information (HTML5)

	element := &Element{
		tagName:    "details",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Pre(elements ...*Element) *Element {
	// Preformatted text

	element := &Element{
		tagName:    "pre",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Bdi(elements ...*Element) *Element {
	// BiDi isolate (HTML5)

	element := &Element{
		tagName:    "bdi",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H3(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h3",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H6(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h6",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Nav(elements ...*Element) *Element {
	// Group of navigational links (HTML5)

	element := &Element{
		tagName:    "nav",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Header(elements ...*Element) *Element {
	// Header (HTML5)

	element := &Element{
		tagName:    "header",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Optgroup(elements ...*Element) *Element {
	// Group of options

	element := &Element{
		tagName:    "optgroup",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Param(elements ...*Element) *Element {
	// Initialization parameters for plugins

	element := &Element{
		tagName:    "param",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Summary(elements ...*Element) *Element {
	// Summary, caption, or legend for a details control (HTML5)

	element := &Element{
		tagName:    "summary",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Fieldset(elements ...*Element) *Element {
	// Set of related form controls

	element := &Element{
		tagName:    "fieldset",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H4(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h4",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Html(elements ...*Element) *Element {
	// Root element

	element := &Element{
		tagName:    "html",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Sub(elements ...*Element) *Element {
	// Subscript

	element := &Element{
		tagName:    "sub",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Track(elements ...*Element) *Element {
	// Supplementary media track (HTML5)

	element := &Element{
		tagName:    "track",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Embed(elements ...*Element) *Element {
	// Integration point for plugins (HTML5)

	element := &Element{
		tagName:    "embed",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Abbr(elements ...*Element) *Element {
	// Abbreviation

	element := &Element{
		tagName:    "abbr",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Body(elements ...*Element) *Element {
	// Document body

	element := &Element{
		tagName:    "body",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Ol(elements ...*Element) *Element {
	// Ordered list

	element := &Element{
		tagName:    "ol",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Output(elements ...*Element) *Element {
	// Result of a calculation in a form (HTML5)

	element := &Element{
		tagName:    "output",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Small(elements ...*Element) *Element {
	// Small print

	element := &Element{
		tagName:    "small",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H1(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h1",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Head(elements ...*Element) *Element {
	// Document metadata container

	element := &Element{
		tagName:    "head",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Link(elements ...*Element) *Element {
	// Inter-document relationship metadata

	element := &Element{
		tagName:    "link",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Time(elements ...*Element) *Element {
	// Date and/or time (HTML5)

	element := &Element{
		tagName:    "time",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Ul(elements ...*Element) *Element {
	// Unordered list

	element := &Element{
		tagName:    "ul",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func B(elements ...*Element) *Element {
	// Offset text conventionally styled in bold

	element := &Element{
		tagName:    "b",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Progress(elements ...*Element) *Element {
	// Progress indicator (HTML5)

	element := &Element{
		tagName:    "progress",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Footer(elements ...*Element) *Element {
	// Footer (HTML5)

	element := &Element{
		tagName:    "footer",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Noscript(elements ...*Element) *Element {
	// Fallback content for script

	element := &Element{
		tagName:    "noscript",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Title(elements ...*Element) *Element {
	// Document title

	element := &Element{
		tagName:    "title",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func P(elements ...*Element) *Element {
	// Paragraph

	element := &Element{
		tagName:    "p",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Style(elements ...*Element) *Element {
	// Style (presentation) information

	element := &Element{
		tagName:    "style",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Tfoot(elements ...*Element) *Element {
	// Table footer row group

	element := &Element{
		tagName:    "tfoot",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Video(elements ...*Element) *Element {
	// Video (HTML5)

	element := &Element{
		tagName:    "video",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Audio(elements ...*Element) *Element {
	// Audio stream (HTML5)

	element := &Element{
		tagName:    "audio",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Datalist(elements ...*Element) *Element {
	// Predefined options for other controls (HTML5)

	element := &Element{
		tagName:    "datalist",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Select(elements ...*Element) *Element {
	// Option-selection form control

	element := &Element{
		tagName:    "select",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Dt(elements ...*Element) *Element {
	// Term or name

	element := &Element{
		tagName:    "dt",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Hgroup(elements ...*Element) *Element {
	// Heading group (HTML5)

	element := &Element{
		tagName:    "hgroup",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Label(elements ...*Element) *Element {
	// Caption for a form control

	element := &Element{
		tagName:    "label",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Menu(elements ...*Element) *Element {
	// List of commands

	element := &Element{
		tagName:    "menu",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func S(elements ...*Element) *Element {
	// Struck text

	element := &Element{
		tagName:    "s",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Table(elements ...*Element) *Element {
	// Table

	element := &Element{
		tagName:    "table",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Cite(elements ...*Element) *Element {
	// Cited title of a work

	element := &Element{
		tagName:    "cite",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Mark(elements ...*Element) *Element {
	// Marked (highlighted) text (HTML5)

	element := &Element{
		tagName:    "mark",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Samp(elements ...*Element) *Element {
	// (sample) output

	element := &Element{
		tagName:    "samp",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Script(elements ...*Element) *Element {
	// Embedded script

	element := &Element{
		tagName:    "script",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Tbody(elements ...*Element) *Element {
	// Table row group

	element := &Element{
		tagName:    "tbody",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Canvas(elements ...*Element) *Element {
	// Canvas for dynamic graphics (HTML5)

	element := &Element{
		tagName:    "canvas",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Dfn(elements ...*Element) *Element {
	// Defining instance

	element := &Element{
		tagName:    "dfn",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Dl(elements ...*Element) *Element {
	// Description list

	element := &Element{
		tagName:    "dl",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Li(elements ...*Element) *Element {
	// List item

	element := &Element{
		tagName:    "li",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Q(elements ...*Element) *Element {
	// Quoted text

	element := &Element{
		tagName:    "q",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Bdo(elements ...*Element) *Element {
	// BiDi override

	element := &Element{
		tagName:    "bdo",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Tr(elements ...*Element) *Element {
	// Table row

	element := &Element{
		tagName:    "tr",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Br(elements ...*Element) *Element {
	// Line break

	element := &Element{
		tagName:    "br",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Kbd(elements ...*Element) *Element {
	// User input

	element := &Element{
		tagName:    "kbd",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Option(elements ...*Element) *Element {
	// Option

	element := &Element{
		tagName:    "option",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Span(elements ...*Element) *Element {
	// Generic span

	element := &Element{
		tagName:    "span",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Th(elements ...*Element) *Element {
	// Table header cell

	element := &Element{
		tagName:    "th",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Var(elements ...*Element) *Element {
	// Variable or placeholder text

	element := &Element{
		tagName:    "var",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func A(elements ...*Element) *Element {
	// Hyperlink

	element := &Element{
		tagName:    "a",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func H2(elements ...*Element) *Element {
	// Heading

	element := &Element{
		tagName:    "h2",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Ins(elements ...*Element) *Element {
	// Inserted text

	element := &Element{
		tagName:    "ins",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Map(elements ...*Element) *Element {
	// Image-map definition

	element := &Element{
		tagName:    "map",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}

func Thead(elements ...*Element) *Element {
	// Table heading group

	element := &Element{
		tagName:    "thead",
		attributes: make(map[string]string),
		children:   elements,
	}
	for _, child := range elements {
		child.parent = element
	}
	return element
}
