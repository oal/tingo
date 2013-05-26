package tingo

// Image-map hyperlink
func Area() *Element {
	return newVoidElement("area")
}

// Result of a calculation in a form (HTML5)
func Output(children ...*Element) *Element {
	return newElement("output", children)
}

// Initialization parameters for plugins
func Param() *Element {
	return newVoidElement("param")
}

// Embedded script
func Script(children ...*Element) *Element {
	return newElement("script", children)
}

// Figure caption (HTML5)
func Figcaption(children ...*Element) *Element {
	return newElement("figcaption", children)
}

// Caption for a form control
func Label(children ...*Element) *Element {
	return newElement("label", children)
}

// Small print
func Small(children ...*Element) *Element {
	return newElement("small", children)
}

// Predefined options for other controls (HTML5)
func Datalist(children ...*Element) *Element {
	return newElement("datalist", children)
}

// User-submittable form
func Form(children ...*Element) *Element {
	return newElement("form", children)
}

// Inserted text
func Ins(children ...*Element) *Element {
	return newElement("ins", children)
}

// Text input area
func Textarea(children ...*Element) *Element {
	return newElement("textarea", children)
}

// Line break
func Br() *Element {
	return newVoidElement("br")
}

// Style (presentation) information
func Style(children ...*Element) *Element {
	return newElement("style", children)
}

// Date and/or time (HTML5)
func Time(children ...*Element) *Element {
	return newElement("time", children)
}

// Paragraph
func P(children ...*Element) *Element {
	return newElement("p", children)
}

// Offset text conventionally styled with an underline
func U(children ...*Element) *Element {
	return newElement("u", children)
}

// Variable or placeholder text
func Var(children ...*Element) *Element {
	return newElement("var", children)
}

// Article (HTML5)
func Article(children ...*Element) *Element {
	return newElement("article", children)
}

// Tangential content (HTML5)
func Aside(children ...*Element) *Element {
	return newElement("aside", children)
}

// Table column
func Col() *Element {
	return newVoidElement("col")
}

// Figure with optional caption (HTML5)
func Figure(children ...*Element) *Element {
	return newElement("figure", children)
}

// Heading group (HTML5)
func Hgroup(children ...*Element) *Element {
	return newElement("hgroup", children)
}

// Quoted text
func Q(children ...*Element) *Element {
	return newElement("q", children)
}

// Unordered list
func Ul(children ...*Element) *Element {
	return newElement("ul", children)
}

// Cited title of a work
func Cite(children ...*Element) *Element {
	return newElement("cite", children)
}

// Document metadata container
func Head(children ...*Element) *Element {
	return newElement("head", children)
}

// Strong importance
func Strong(children ...*Element) *Element {
	return newElement("strong", children)
}

// Subscript
func Sub(children ...*Element) *Element {
	return newElement("sub", children)
}

// Table column group
func Colgroup(children ...*Element) *Element {
	return newElement("colgroup", children)
}

// Metadata
func Meta() *Element {
	return newVoidElement("meta")
}

// Generic span
func Span(children ...*Element) *Element {
	return newElement("span", children)
}

// Key-pair generator/input control (HTML5)
func Keygen() *Element {
	return newVoidElement("keygen")
}

// Line-break opportunity (HTML5)
func Wbr() *Element {
	return newVoidElement("wbr")
}

// Code fragment
func Code(children ...*Element) *Element {
	return newElement("code", children)
}

// Heading
func H3(children ...*Element) *Element {
	return newElement("h3", children)
}

// Thematic break
func Hr() *Element {
	return newVoidElement("hr")
}

// Generic external content
func Object(children ...*Element) *Element {
	return newElement("object", children)
}

// Struck text
func S(children ...*Element) *Element {
	return newElement("s", children)
}

// Summary, caption, or legend for a details control (HTML5)
func Summary(children ...*Element) *Element {
	return newElement("summary", children)
}

// Audio stream (HTML5)
func Audio(children ...*Element) *Element {
	return newElement("audio", children)
}

// Description or value
func Dd(children ...*Element) *Element {
	return newElement("dd", children)
}

// Emphatic stress
func Em(children ...*Element) *Element {
	return newElement("em", children)
}

// Heading
func H6(children ...*Element) *Element {
	return newElement("h6", children)
}

// Input control
func Input() *Element {
	return newVoidElement("input")
}

// Generic function to create the various inputs
func genericInput(name, t string) *Element {
	el := newVoidElement("input")
	el.attributes["name"] = name
	el.attributes["type"] = t
	return el
}

// Text input control
func TextInput(name string) *Element {
	return genericInput(name, "text")
}

// Password input control
func PasswordInput(name string) *Element {
	return genericInput(name, "password")
}

// Checkbox input control
func CheckboxInput(name string) *Element {
	return genericInput(name, "checkbox")
}

// Radio input control
func RadioInput(name string) *Element {
	return genericInput(name, "radio")
}

// Button input control
func ButtonInput(name string) *Element {
	return genericInput(name, "button")
}

// Submit input control
func SubmitInput(name string) *Element {
	return genericInput(name, "submit")
}

// Reset input control
func ResetInput(name string) *Element {
	return genericInput(name, "reset")
}

// File input control
func FileInput(name string) *Element {
	return genericInput(name, "file")
}

// Hidden input control
func HiddenInput(name string) *Element {
	return genericInput(name, "hidden")
}

// Datetime input control (HTML5)
func DatetimeInput(name string) *Element {
	return genericInput(name, "datetime")
}

// Datetime local input control (HTML5)
func DatetimeLocalInput(name string) *Element {
	return genericInput(name, "datetime-local")
}

// Date input control (HTML5)
func DateInput(name string) *Element {
	return genericInput(name, "date")
}

// Month input control (HTML5)
func MonthInput(name string) *Element {
	return genericInput(name, "month")
}

// Time input control (HTML5)
func TimeInput(name string) *Element {
	return genericInput(name, "time")
}

// Week input control (HTML5)
func WeekInput(name string) *Element {
	return genericInput(name, "week")
}

// Number input control (HTML5)
func NumberInput(name string) *Element {
	return genericInput(name, "number")
}

// Range input control (HTML5)
func RangeInput(name string) *Element {
	return genericInput(name, "range")
}

// Email input control (HTML5)
func EmailInput(name string) *Element {
	return genericInput(name, "email")
}

// Url input control (HTML5)
func UrlInput(name string) *Element {
	return genericInput(name, "url")
}

// Search input control (HTML5)
func SearchInput(name string) *Element {
	return genericInput(name, "search")
}

// Tel input control (HTML5)
func TelInput(name string) *Element {
	return genericInput(name, "tel")
}

// Color input control (HTML5)
func ColorInput(name string) *Element {
	return genericInput(name, "color")
}

// Marked (highlighted) text (HTML5)
func Mark(children ...*Element) *Element {
	return newElement("mark", children)
}

// Ruby text (HTML5)
func Rt(children ...*Element) *Element {
	return newElement("rt", children)
}

// Table heading group
func Thead(children ...*Element) *Element {
	return newElement("thead", children)
}

// User input
func Kbd(children ...*Element) *Element {
	return newElement("kbd", children)
}

// Ruby parenthesis (HTML5)
func Rp(children ...*Element) *Element {
	return newElement("rp", children)
}

// Offset text conventionally styled in bold
func B(children ...*Element) *Element {
	return newElement("b", children)
}

// BiDi override
func Bdo(children ...*Element) *Element {
	return newElement("bdo", children)
}

// Button
func Button(children ...*Element) *Element {
	return newElement("button", children)
}

// Fallback content for script
func Noscript(children ...*Element) *Element {
	return newElement("noscript", children)
}

// Document title
func Title(children ...*Element) *Element {
	return newElement("title", children)
}

// Hyperlink
func A(children ...*Element) *Element {
	return newElement("a", children)
}

// Control for additional on-demand information (HTML5)
func Details(children ...*Element) *Element {
	return newElement("details", children)
}

// Integration point for plugins (HTML5)
func Embed() *Element {
	return newVoidElement("embed")
}

// Option-selection form control
func Select(children ...*Element) *Element {
	return newElement("select", children)
}

// Footer (HTML5)
func Footer(children ...*Element) *Element {
	return newElement("footer", children)
}

// Table
func Table(children ...*Element) *Element {
	return newElement("table", children)
}

// Offset text conventionally styled in italic
func I(children ...*Element) *Element {
	return newElement("i", children)
}

// List of commands
func Menu(children ...*Element) *Element {
	return newElement("menu", children)
}

// (sample) output
func Samp(children ...*Element) *Element {
	return newElement("samp", children)
}

// Block quotation
func Blockquote(children ...*Element) *Element {
	return newElement("blockquote", children)
}

// Media source (HTML5)
func Source() *Element {
	return newVoidElement("source")
}

// Table row group
func Tbody(children ...*Element) *Element {
	return newElement("tbody", children)
}

// Heading
func H2(children ...*Element) *Element {
	return newElement("h2", children)
}

// Image-map definition
func Map(children ...*Element) *Element {
	return newElement("map", children)
}

// Term or name
func Dt(children ...*Element) *Element {
	return newElement("dt", children)
}

// Header (HTML5)
func Header(children ...*Element) *Element {
	return newElement("header", children)
}

// Ordered list
func Ol(children ...*Element) *Element {
	return newElement("ol", children)
}

// Group of options
func Optgroup(children ...*Element) *Element {
	return newElement("optgroup", children)
}

// Description list
func Dl(children ...*Element) *Element {
	return newElement("dl", children)
}

// Title or explanatory caption
func Legend(children ...*Element) *Element {
	return newElement("legend", children)
}

// Progress indicator (HTML5)
func Progress(children ...*Element) *Element {
	return newElement("progress", children)
}

// Table header cell
func Th(children ...*Element) *Element {
	return newElement("th", children)
}

// Table row
func Tr(children ...*Element) *Element {
	return newElement("tr", children)
}

// Scalar gauge (HTML5)
func Meter(children ...*Element) *Element {
	return newElement("meter", children)
}

// Ruby annotation (HTML5)
func Ruby(children ...*Element) *Element {
	return newElement("ruby", children)
}

// Set of related form controls
func Fieldset(children ...*Element) *Element {
	return newElement("fieldset", children)
}

// Heading
func H1(children ...*Element) *Element {
	return newElement("h1", children)
}

// Superscript
func Sup(children ...*Element) *Element {
	return newElement("sup", children)
}

// Table footer row group
func Tfoot(children ...*Element) *Element {
	return newElement("tfoot", children)
}

// Supplementary media track (HTML5)
func Track() *Element {
	return newVoidElement("track")
}

// Abbreviation
func Abbr(children ...*Element) *Element {
	return newElement("abbr", children)
}

// Document body
func Body(children ...*Element) *Element {
	return newElement("body", children)
}

// Command (HTML5)
func Command() *Element {
	return newVoidElement("command")
}

// Generic flow container
func Div(children ...*Element) *Element {
	return newElement("div", children)
}

// Nested browsing context (inline frame)
func Iframe(children ...*Element) *Element {
	return newElement("iframe", children)
}

// Image
func Img() *Element {
	return newVoidElement("img")
}

// List item
func Li(children ...*Element) *Element {
	return newElement("li", children)
}

// Inter-document relationship metadata
func Link() *Element {
	return newVoidElement("link")
}

// Section (HTML5)
func Section(children ...*Element) *Element {
	return newElement("section", children)
}

// Heading
func H4(children ...*Element) *Element {
	return newElement("h4", children)
}

// BiDi isolate (HTML5)
func Bdi(children ...*Element) *Element {
	return newElement("bdi", children)
}

// Defining instance
func Dfn(children ...*Element) *Element {
	return newElement("dfn", children)
}

// Root element
func Html(children ...*Element) *Element {
	return newElement("html", children)
}

// Option
func Option(children ...*Element) *Element {
	return newElement("option", children)
}

// Table title
func Caption(children ...*Element) *Element {
	return newElement("caption", children)
}

// Group of navigational links (HTML5)
func Nav(children ...*Element) *Element {
	return newElement("nav", children)
}

// Table cell
func Td(children ...*Element) *Element {
	return newElement("td", children)
}

// Base URL
func Base() *Element {
	return newVoidElement("base")
}

// Canvas for dynamic graphics (HTML5)
func Canvas(children ...*Element) *Element {
	return newElement("canvas", children)
}

// Deleted text
func Del(children ...*Element) *Element {
	return newElement("del", children)
}

// Heading
func H5(children ...*Element) *Element {
	return newElement("h5", children)
}

// Video (HTML5)
func Video(children ...*Element) *Element {
	return newElement("video", children)
}

// Contact information
func Address(children ...*Element) *Element {
	return newElement("address", children)
}

// Preformatted text
func Pre(children ...*Element) *Element {
	return newElement("pre", children)
}
