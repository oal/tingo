package tingo

func Li(children ...*Element) *Element {
	// List item
	return newElement("li", children)
}

func Video(children ...*Element) *Element {
	// Video (HTML5)
	return newElement("video", children)
}

func Embed() *Element {
	// Integration point for plugins (HTML5)
	return newVoidElement("embed")
}

func Input() *Element {
	// Input control
	return newVoidElement("input")
}

func Nav(children ...*Element) *Element {
	// Group of navigational links (HTML5)
	return newElement("nav", children)
}

func Option(children ...*Element) *Element {
	// Option
	return newElement("option", children)
}

func Output(children ...*Element) *Element {
	// Result of a calculation in a form (HTML5)
	return newElement("output", children)
}

func Var(children ...*Element) *Element {
	// Variable or placeholder text
	return newElement("var", children)
}

func Label(children ...*Element) *Element {
	// Caption for a form control
	return newElement("label", children)
}

func Object(children ...*Element) *Element {
	// Generic external content
	return newElement("object", children)
}

func Rt(children ...*Element) *Element {
	// Ruby text (HTML5)
	return newElement("rt", children)
}

func Dfn(children ...*Element) *Element {
	// Defining instance
	return newElement("dfn", children)
}

func Header(children ...*Element) *Element {
	// Header (HTML5)
	return newElement("header", children)
}

func Progress(children ...*Element) *Element {
	// Progress indicator (HTML5)
	return newElement("progress", children)
}

func Base() *Element {
	// Base URL
	return newVoidElement("base")
}

func Form(children ...*Element) *Element {
	// User-submittable form
	return newElement("form", children)
}

func Aside(children ...*Element) *Element {
	// Tangential content (HTML5)
	return newElement("aside", children)
}

func Hgroup(children ...*Element) *Element {
	// Heading group (HTML5)
	return newElement("hgroup", children)
}

func Ol(children ...*Element) *Element {
	// Ordered list
	return newElement("ol", children)
}

func Table(children ...*Element) *Element {
	// Table
	return newElement("table", children)
}

func Blockquote(children ...*Element) *Element {
	// Block quotation
	return newElement("blockquote", children)
}

func Datalist(children ...*Element) *Element {
	// Predefined options for other controls (HTML5)
	return newElement("datalist", children)
}

func I(children ...*Element) *Element {
	// Offset text conventionally styled in italic
	return newElement("i", children)
}

func Link() *Element {
	// Inter-document relationship metadata
	return newVoidElement("link")
}

func Thead(children ...*Element) *Element {
	// Table heading group
	return newElement("thead", children)
}

func A(children ...*Element) *Element {
	// Hyperlink
	return newElement("a", children)
}

func Noscript(children ...*Element) *Element {
	// Fallback content for script
	return newElement("noscript", children)
}

func Td(children ...*Element) *Element {
	// Table cell
	return newElement("td", children)
}

func H1(children ...*Element) *Element {
	// Heading
	return newElement("h1", children)
}

func Img() *Element {
	// Image
	return newVoidElement("img")
}

func Textarea(children ...*Element) *Element {
	// Text input area
	return newElement("textarea", children)
}

func Body(children ...*Element) *Element {
	// Document body
	return newElement("body", children)
}

func Em(children ...*Element) *Element {
	// Emphatic stress
	return newElement("em", children)
}

func Mark(children ...*Element) *Element {
	// Marked (highlighted) text (HTML5)
	return newElement("mark", children)
}

func Track() *Element {
	// Supplementary media track (HTML5)
	return newVoidElement("track")
}

func Address(children ...*Element) *Element {
	// Contact information
	return newElement("address", children)
}

func H3(children ...*Element) *Element {
	// Heading
	return newElement("h3", children)
}

func H6(children ...*Element) *Element {
	// Heading
	return newElement("h6", children)
}

func Kbd(children ...*Element) *Element {
	// User input
	return newElement("kbd", children)
}

func Rp(children ...*Element) *Element {
	// Ruby parenthesis (HTML5)
	return newElement("rp", children)
}

func Wbr() *Element {
	// Line-break opportunity (HTML5)
	return newVoidElement("wbr")
}

func Dt(children ...*Element) *Element {
	// Term or name
	return newElement("dt", children)
}

func Meta() *Element {
	// Metadata
	return newVoidElement("meta")
}

func Tbody(children ...*Element) *Element {
	// Table row group
	return newElement("tbody", children)
}

func H5(children ...*Element) *Element {
	// Heading
	return newElement("h5", children)
}

func Tr(children ...*Element) *Element {
	// Table row
	return newElement("tr", children)
}

func Audio(children ...*Element) *Element {
	// Audio stream (HTML5)
	return newElement("audio", children)
}

func Ins(children ...*Element) *Element {
	// Inserted text
	return newElement("ins", children)
}

func Legend(children ...*Element) *Element {
	// Title or explanatory caption
	return newElement("legend", children)
}

func Del(children ...*Element) *Element {
	// Deleted text
	return newElement("del", children)
}

func Iframe(children ...*Element) *Element {
	// Nested browsing context (inline frame)
	return newElement("iframe", children)
}

func Pre(children ...*Element) *Element {
	// Preformatted text
	return newElement("pre", children)
}

func Script(children ...*Element) *Element {
	// Embedded script
	return newElement("script", children)
}

func Select(children ...*Element) *Element {
	// Option-selection form control
	return newElement("select", children)
}

func Head(children ...*Element) *Element {
	// Document metadata container
	return newElement("head", children)
}

func Html(children ...*Element) *Element {
	// Root element
	return newElement("html", children)
}

func Abbr(children ...*Element) *Element {
	// Abbreviation
	return newElement("abbr", children)
}

func Div(children ...*Element) *Element {
	// Generic flow container
	return newElement("div", children)
}

func Style(children ...*Element) *Element {
	// Style (presentation) information
	return newElement("style", children)
}

func Sup(children ...*Element) *Element {
	// Superscript
	return newElement("sup", children)
}

func Tfoot(children ...*Element) *Element {
	// Table footer row group
	return newElement("tfoot", children)
}

func Canvas(children ...*Element) *Element {
	// Canvas for dynamic graphics (HTML5)
	return newElement("canvas", children)
}

func Figure(children ...*Element) *Element {
	// Figure with optional caption (HTML5)
	return newElement("figure", children)
}

func Optgroup(children ...*Element) *Element {
	// Group of options
	return newElement("optgroup", children)
}

func Strong(children ...*Element) *Element {
	// Strong importance
	return newElement("strong", children)
}

func Menu(children ...*Element) *Element {
	// List of commands
	return newElement("menu", children)
}

func S(children ...*Element) *Element {
	// Struck text
	return newElement("s", children)
}

func Bdo(children ...*Element) *Element {
	// BiDi override
	return newElement("bdo", children)
}

func Fieldset(children ...*Element) *Element {
	// Set of related form controls
	return newElement("fieldset", children)
}

func H4(children ...*Element) *Element {
	// Heading
	return newElement("h4", children)
}

func Param() *Element {
	// Initialization parameters for plugins
	return newVoidElement("param")
}

func Time(children ...*Element) *Element {
	// Date and/or time (HTML5)
	return newElement("time", children)
}

func U(children ...*Element) *Element {
	// Offset text conventionally styled with an underline
	return newElement("u", children)
}

func Hr() *Element {
	// Thematic break
	return newVoidElement("hr")
}

func P(children ...*Element) *Element {
	// Paragraph
	return newElement("p", children)
}

func Title(children ...*Element) *Element {
	// Document title
	return newElement("title", children)
}

func Ul(children ...*Element) *Element {
	// Unordered list
	return newElement("ul", children)
}

func Figcaption(children ...*Element) *Element {
	// Figure caption (HTML5)
	return newElement("figcaption", children)
}

func B(children ...*Element) *Element {
	// Offset text conventionally styled in bold
	return newElement("b", children)
}

func Bdi(children ...*Element) *Element {
	// BiDi isolate (HTML5)
	return newElement("bdi", children)
}

func Cite(children ...*Element) *Element {
	// Cited title of a work
	return newElement("cite", children)
}

func Caption(children ...*Element) *Element {
	// Table title
	return newElement("caption", children)
}

func Col() *Element {
	// Table column
	return newVoidElement("col")
}

func Meter(children ...*Element) *Element {
	// Scalar gauge (HTML5)
	return newElement("meter", children)
}

func Summary(children ...*Element) *Element {
	// Summary, caption, or legend for a details control (HTML5)
	return newElement("summary", children)
}

func Colgroup(children ...*Element) *Element {
	// Table column group
	return newElement("colgroup", children)
}

func Samp(children ...*Element) *Element {
	// (sample) output
	return newElement("samp", children)
}

func Span(children ...*Element) *Element {
	// Generic span
	return newElement("span", children)
}

func Sub(children ...*Element) *Element {
	// Subscript
	return newElement("sub", children)
}

func Article(children ...*Element) *Element {
	// Article (HTML5)
	return newElement("article", children)
}

func Code(children ...*Element) *Element {
	// Code fragment
	return newElement("code", children)
}

func Map(children ...*Element) *Element {
	// Image-map definition
	return newElement("map", children)
}

func Ruby(children ...*Element) *Element {
	// Ruby annotation (HTML5)
	return newElement("ruby", children)
}

func Dd(children ...*Element) *Element {
	// Description or value
	return newElement("dd", children)
}

func Keygen() *Element {
	// Key-pair generator/input control (HTML5)
	return newVoidElement("keygen")
}

func Section(children ...*Element) *Element {
	// Section (HTML5)
	return newElement("section", children)
}

func Source() *Element {
	// Media source (HTML5)
	return newVoidElement("source")
}

func Br() *Element {
	// Line break
	return newVoidElement("br")
}

func Command() *Element {
	// Command (HTML5)
	return newVoidElement("command")
}

func Small(children ...*Element) *Element {
	// Small print
	return newElement("small", children)
}

func Button(children ...*Element) *Element {
	// Button
	return newElement("button", children)
}

func Footer(children ...*Element) *Element {
	// Footer (HTML5)
	return newElement("footer", children)
}

func Dl(children ...*Element) *Element {
	// Description list
	return newElement("dl", children)
}

func H2(children ...*Element) *Element {
	// Heading
	return newElement("h2", children)
}

func Th(children ...*Element) *Element {
	// Table header cell
	return newElement("th", children)
}

func Area() *Element {
	// Image-map hyperlink
	return newVoidElement("area")
}

func Details(children ...*Element) *Element {
	// Control for additional on-demand information (HTML5)
	return newElement("details", children)
}

func Q(children ...*Element) *Element {
	// Quoted text
	return newElement("q", children)
}
