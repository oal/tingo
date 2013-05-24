package tingo

/*

Document related

*/

type HtmlElement struct{ *Element }
type HeadElement struct{ *Element }
type MetaElement struct{ *Element }
type StyleElement struct{ *Element }
type BodyElement struct{ *Element }

func Html(children ...ElementInterface) *HtmlElement {
	// Root element
	return &HtmlElement{newElement("html", children)}
}

func Head(children ...ElementInterface) *HeadElement {
	// Document metadata container
	return &HeadElement{newElement("head", children)}
}

func Meta(children ...ElementInterface) *MetaElement {
	// Metadata
	return &MetaElement{newElement("meta", children)}
}

func Style(children ...ElementInterface) *StyleElement {
	// Style (presentation) information
	return &StyleElement{newElement("style", children)}
}

func Body(children ...ElementInterface) *BodyElement {
	// Document body
	return &BodyElement{newElement("body", children)}
}

/*

Formatting

*/

type BElement struct{ *Element }
type IElement struct{ *Element }

func B(children ...ElementInterface) *BElement {
	// Offset text conventionally styled in bold
	return &BElement{newElement("b", children)}
}

func I(children ...ElementInterface) *IElement {
	// Offset text conventionally styled in italic
	return &IElement{newElement("i", children)}
}

type StrongElement struct{ *Element }
type EmElement struct{ *Element }
type SmallElement struct{ *Element }

func Strong(children ...ElementInterface) *StrongElement {
	// Strong importance
	return &StrongElement{newElement("strong", children)}
}

func Em(children ...ElementInterface) *EmElement {
	// Emphatic stress
	return &EmElement{newElement("em", children)}
}

func Small(children ...ElementInterface) *SmallElement {
	// Small print
	return &SmallElement{newElement("small", children)}
}

type SupElement struct{ *Element }
type SubElement struct{ *Element }

func Sup(children ...ElementInterface) *SupElement {
	// Superscript
	return &SupElement{newElement("sup", children)}
}

func Sub(children ...ElementInterface) *SubElement {
	// Subscript
	return &SubElement{newElement("sub", children)}
}

type BrElement struct{ *Element }
type HrElement struct{ *Element }

func Br(children ...ElementInterface) *BrElement {
	// Line break
	return &BrElement{newElement("br", children)}
}

func Hr(children ...ElementInterface) *HrElement {
	// Thematic break
	return &HrElement{newElement("hr", children)}
}

type CodeElement struct{ *Element }
type BlockquoteElement struct{ *Element }

func Code(children ...ElementInterface) *CodeElement {
	// Code fragment
	return &CodeElement{newElement("code", children)}
}

func Blockquote(children ...ElementInterface) *BlockquoteElement {
	// Block quotation
	return &BlockquoteElement{newElement("blockquote", children)}
}

/*

Lists

*/

type UlElement struct{ *Element }
type OlElement struct{ *Element }
type LiElement struct{ *Element }

func Ul(children ...ElementInterface) *UlElement {
	// Unordered list
	return &UlElement{newElement("ul", children)}
}

func Ol(children ...ElementInterface) *OlElement {
	// Ordered list
	return &OlElement{newElement("ol", children)}
}

func Li(children ...ElementInterface) *LiElement {
	// List item
	return &LiElement{newElement("li", children)}
}

type DlElement struct{ *Element }
type DdElement struct{ *Element }

func Dl(children ...ElementInterface) *DlElement {
	// Description list
	return &DlElement{newElement("dl", children)}
}

func Dd(children ...ElementInterface) *DdElement {
	// Description or value
	return &DdElement{newElement("dd", children)}
}

/*

Table related

*/

type TableElement struct{ *Element }

func Table(children ...ElementInterface) *TableElement {
	// Table
	return &TableElement{newElement("table", children)}
}

type TheadElement struct{ *Element }
type ThElement struct{ *Element }

func Thead(children ...ElementInterface) *TheadElement {
	// Table heading group
	return &TheadElement{newElement("thead", children)}
}

func Th(children ...ElementInterface) *ThElement {
	// Table header cell
	return &ThElement{newElement("th", children)}
}

type TbodyElement struct{ *Element }
type TrElement struct{ *Element }
type TdElement struct{ *Element }

func Tbody(children ...ElementInterface) *TbodyElement {
	// Table row group
	return &TbodyElement{newElement("tbody", children)}
}

func Tr(children ...ElementInterface) *TrElement {
	// Table row
	return &TrElement{newElement("tr", children)}
}

func Td(children ...ElementInterface) *TdElement {
	// Table cell
	return &TdElement{newElement("td", children)}
}

type TfootElement struct{ *Element }

func Tfoot(children ...ElementInterface) *TfootElement {
	// Table footer row group
	return &TfootElement{newElement("tfoot", children)}
}

type ColgroupElement struct{ *Element }
type ColElement struct{ *Element }

func Colgroup(children ...ElementInterface) *ColgroupElement {
	// Table column group
	return &ColgroupElement{newElement("colgroup", children)}
}
func Col(children ...ElementInterface) *ColElement {
	// Table column
	return &ColElement{newElement("col", children)}
}

/*

Forms

*/

type FormElement struct{ *Element }
type InputElement struct{ *Element }
type ButtonElement struct{ *Element }
type TextareaElement struct{ *Element }
type SelectElement struct{ *Element }
type OptionElement struct{ *Element }
type OptgroupElement struct{ *Element }
type LabelElement struct{ *Element }
type FieldsetElement struct{ *Element }

func Form(children ...ElementInterface) *FormElement {
	// User-submittable form
	return &FormElement{newElement("form", children)}
}

func Input(children ...ElementInterface) *InputElement {
	// Input control
	return &InputElement{newElement("input", children)}
}

func (i *InputElement) Type(t string) *InputElement {
	i.attributes["type"] = t
	return i
}

func Button(children ...ElementInterface) *ButtonElement {
	// Button
	return &ButtonElement{newElement("button", children)}
}

func Textarea(children ...ElementInterface) *TextareaElement {
	// Text input area
	return &TextareaElement{newElement("textarea", children)}
}

func Select(children ...ElementInterface) *SelectElement {
	// Option-selection form control
	return &SelectElement{newElement("select", children)}
}

func Option(children ...ElementInterface) *OptionElement {
	// Option
	return &OptionElement{newElement("option", children)}
}

func Optgroup(children ...ElementInterface) *OptgroupElement {
	// Group of options
	return &OptgroupElement{newElement("optgroup", children)}
}

func Label(children ...ElementInterface) *LabelElement {
	// Caption for a form control
	return &LabelElement{newElement("label", children)}
}

func Fieldset(children ...ElementInterface) *FieldsetElement {
	// Set of related form controls
	return &FieldsetElement{newElement("fieldset", children)}
}

/*

Headings

*/

type H1Element struct{ *Element }
type H2Element struct{ *Element }
type H3Element struct{ *Element }
type H4Element struct{ *Element }
type H5Element struct{ *Element }
type H6Element struct{ *Element }
type HgroupElement struct{ *Element }

func H1(children ...ElementInterface) *H1Element {
	// Heading
	return &H1Element{newElement("h1", children)}
}

func H2(children ...ElementInterface) *H2Element {
	// Heading
	return &H2Element{newElement("h2", children)}
}

func H3(children ...ElementInterface) *H3Element {
	// Heading
	return &H3Element{newElement("h3", children)}
}

func H4(children ...ElementInterface) *H4Element {
	// Heading
	return &H4Element{newElement("h4", children)}
}

func H5(children ...ElementInterface) *H5Element {
	// Heading
	return &H5Element{newElement("h5", children)}
}

func H6(children ...ElementInterface) *H6Element {
	// Heading
	return &H6Element{newElement("h6", children)}
}

func Hgroup(children ...ElementInterface) *HgroupElement {
	// Heading group (HTML5)
	return &HgroupElement{newElement("hgroup", children)}
}

/*

TODO: Organize the elements below

*/

type VarElement struct{ *Element }

func Var(children ...ElementInterface) *VarElement {
	// Variable or placeholder text
	return &VarElement{newElement("var", children)}
}

type LinkElement struct{ *Element }

func Link(children ...ElementInterface) *LinkElement {
	// Inter-document relationship metadata
	return &LinkElement{newElement("link", children)}
}

type OutputElement struct{ *Element }

func Output(children ...ElementInterface) *OutputElement {
	// Result of a calculation in a form (HTML5)
	return &OutputElement{newElement("output", children)}
}

type AddressElement struct{ *Element }

func Address(children ...ElementInterface) *AddressElement {
	// Contact information
	return &AddressElement{newElement("address", children)}
}

type BdiElement struct{ *Element }

func Bdi(children ...ElementInterface) *BdiElement {
	// BiDi isolate (HTML5)
	return &BdiElement{newElement("bdi", children)}
}

type CiteElement struct{ *Element }

func Cite(children ...ElementInterface) *CiteElement {
	// Cited title of a work
	return &CiteElement{newElement("cite", children)}
}

type VideoElement struct{ *Element }

func Video(children ...ElementInterface) *VideoElement {
	// Video (HTML5)
	return &VideoElement{newElement("video", children)}
}

type IframeElement struct{ *Element }

func Iframe(children ...ElementInterface) *IframeElement {
	// Nested browsing context (inline frame)
	return &IframeElement{newElement("iframe", children)}
}

type ArticleElement struct{ *Element }

func Article(children ...ElementInterface) *ArticleElement {
	// Article (HTML5)
	return &ArticleElement{newElement("article", children)}
}

type CaptionElement struct{ *Element }

func Caption(children ...ElementInterface) *CaptionElement {
	// Table title
	return &CaptionElement{newElement("caption", children)}
}

type ParamElement struct{ *Element }

func Param(children ...ElementInterface) *ParamElement {
	// Initialization parameters for plugins
	return &ParamElement{newElement("param", children)}
}

type AsideElement struct{ *Element }

func Aside(children ...ElementInterface) *AsideElement {
	// Tangential content (HTML5)
	return &AsideElement{newElement("aside", children)}
}

type MapElement struct{ *Element }

func Map(children ...ElementInterface) *MapElement {
	// Image-map definition
	return &MapElement{newElement("map", children)}
}

type FigureElement struct{ *Element }

func Figure(children ...ElementInterface) *FigureElement {
	// Figure with optional caption (HTML5)
	return &FigureElement{newElement("figure", children)}
}

type NavElement struct{ *Element }

func Nav(children ...ElementInterface) *NavElement {
	// Group of navigational links (HTML5)
	return &NavElement{newElement("nav", children)}
}

type RubyElement struct{ *Element }

func Ruby(children ...ElementInterface) *RubyElement {
	// Ruby annotation (HTML5)
	return &RubyElement{newElement("ruby", children)}
}

type EmbedElement struct{ *Element }

func Embed(children ...ElementInterface) *EmbedElement {
	// Integration point for plugins (HTML5)
	return &EmbedElement{newElement("embed", children)}
}

type SampElement struct{ *Element }

func Samp(children ...ElementInterface) *SampElement {
	// (sample) output
	return &SampElement{newElement("samp", children)}
}

type SectionElement struct{ *Element }

func Section(children ...ElementInterface) *SectionElement {
	// Section (HTML5)
	return &SectionElement{newElement("section", children)}
}

type AudioElement struct{ *Element }

func Audio(children ...ElementInterface) *AudioElement {
	// Audio stream (HTML5)
	return &AudioElement{newElement("audio", children)}
}

type DtElement struct{ *Element }

func Dt(children ...ElementInterface) *DtElement {
	// Term or name
	return &DtElement{newElement("dt", children)}
}

type DfnElement struct{ *Element }

func Dfn(children ...ElementInterface) *DfnElement {
	// Defining instance
	return &DfnElement{newElement("dfn", children)}
}

type ObjectElement struct{ *Element }

func Object(children ...ElementInterface) *ObjectElement {
	// Generic external content
	return &ObjectElement{newElement("object", children)}
}

type ProgressElement struct{ *Element }

func Progress(children ...ElementInterface) *ProgressElement {
	// Progress indicator (HTML5)
	return &ProgressElement{newElement("progress", children)}
}

type FigcaptionElement struct{ *Element }

func Figcaption(children ...ElementInterface) *FigcaptionElement {
	// Figure caption (HTML5)
	return &FigcaptionElement{newElement("figcaption", children)}
}

type WbrElement struct{ *Element }

func Wbr(children ...ElementInterface) *WbrElement {
	// Line-break opportunity (HTML5)
	return &WbrElement{newElement("wbr", children)}
}

type CanvasElement struct{ *Element }

func Canvas(children ...ElementInterface) *CanvasElement {
	// Canvas for dynamic graphics (HTML5)
	return &CanvasElement{newElement("canvas", children)}
}

type HeaderElement struct{ *Element }

func Header(children ...ElementInterface) *HeaderElement {
	// Header (HTML5)
	return &HeaderElement{newElement("header", children)}
}

type KeygenElement struct{ *Element }

func Keygen(children ...ElementInterface) *KeygenElement {
	// Key-pair generator/input control (HTML5)
	return &KeygenElement{newElement("keygen", children)}
}

type DelElement struct{ *Element }

func Del(children ...ElementInterface) *DelElement {
	// Deleted text
	return &DelElement{newElement("del", children)}
}

type DetailsElement struct{ *Element }

func Details(children ...ElementInterface) *DetailsElement {
	// Control for additional on-demand information (HTML5)
	return &DetailsElement{newElement("details", children)}
}

type SpanElement struct{ *Element }

func Span(children ...ElementInterface) *SpanElement {
	// Generic span
	return &SpanElement{newElement("span", children)}
}

type MeterElement struct{ *Element }

func Meter(children ...ElementInterface) *MeterElement {
	// Scalar gauge (HTML5)
	return &MeterElement{newElement("meter", children)}
}

type NoscriptElement struct{ *Element }

func Noscript(children ...ElementInterface) *NoscriptElement {
	// Fallback content for script
	return &NoscriptElement{newElement("noscript", children)}
}

type AreaElement struct{ *Element }

func Area(children ...ElementInterface) *AreaElement {
	// Image-map hyperlink
	return &AreaElement{newElement("area", children)}
}

type PElement struct{ *Element }

func P(children ...ElementInterface) *PElement {
	// Paragraph
	return &PElement{newElement("p", children)}
}

type BaseElement struct{ *Element }

func Base(children ...ElementInterface) *BaseElement {
	// Base URL
	return &BaseElement{newElement("base", children)}
}

type DivElement struct{ *Element }

func Div(children ...ElementInterface) *DivElement {
	// Generic flow container
	return &DivElement{newElement("div", children)}
}

type TrackElement struct{ *Element }

func Track(children ...ElementInterface) *TrackElement {
	// Supplementary media track (HTML5)
	return &TrackElement{newElement("track", children)}
}

type DatalistElement struct{ *Element }

func Datalist(children ...ElementInterface) *DatalistElement {
	// Predefined options for other controls (HTML5)
	return &DatalistElement{newElement("datalist", children)}
}

type FooterElement struct{ *Element }

func Footer(children ...ElementInterface) *FooterElement {
	// Footer (HTML5)
	return &FooterElement{newElement("footer", children)}
}

type QElement struct{ *Element }

func Q(children ...ElementInterface) *QElement {
	// Quoted text
	return &QElement{newElement("q", children)}
}

type ImgElement struct{ *Element }

func Img(children ...ElementInterface) *ImgElement {
	// Image
	return &ImgElement{newElement("img", children)}
}

type SourceElement struct{ *Element }

func Source(children ...ElementInterface) *SourceElement {
	// Media source (HTML5)
	return &SourceElement{newElement("source", children)}
}

type MenuElement struct{ *Element }

func Menu(children ...ElementInterface) *MenuElement {
	// List of commands
	return &MenuElement{newElement("menu", children)}
}

type TitleElement struct{ *Element }

func Title(children ...ElementInterface) *TitleElement {
	// Document title
	return &TitleElement{newElement("title", children)}
}

type AElement struct{ *Element }

func A(children ...ElementInterface) *AElement {
	// Hyperlink
	return &AElement{newElement("a", children)}
}

type TimeElement struct{ *Element }

func Time(children ...ElementInterface) *TimeElement {
	// Date and/or time (HTML5)
	return &TimeElement{newElement("time", children)}
}

type MarkElement struct{ *Element }

func Mark(children ...ElementInterface) *MarkElement {
	// Marked (highlighted) text (HTML5)
	return &MarkElement{newElement("mark", children)}
}

type BdoElement struct{ *Element }

func Bdo(children ...ElementInterface) *BdoElement {
	// BiDi override
	return &BdoElement{newElement("bdo", children)}
}

type InsElement struct{ *Element }

func Ins(children ...ElementInterface) *InsElement {
	// Inserted text
	return &InsElement{newElement("ins", children)}
}

type SElement struct{ *Element }

func S(children ...ElementInterface) *SElement {
	// Struck text
	return &SElement{newElement("s", children)}
}

type UElement struct{ *Element }

func U(children ...ElementInterface) *UElement {
	// Offset text conventionally styled with an underline
	return &UElement{newElement("u", children)}
}

type AbbrElement struct{ *Element }

func Abbr(children ...ElementInterface) *AbbrElement {
	// Abbreviation
	return &AbbrElement{newElement("abbr", children)}
}

type KbdElement struct{ *Element }

func Kbd(children ...ElementInterface) *KbdElement {
	// User input
	return &KbdElement{newElement("kbd", children)}
}

type LegendElement struct{ *Element }

func Legend(children ...ElementInterface) *LegendElement {
	// Title or explanatory caption
	return &LegendElement{newElement("legend", children)}
}

type RpElement struct{ *Element }

func Rp(children ...ElementInterface) *RpElement {
	// Ruby parenthesis (HTML5)
	return &RpElement{newElement("rp", children)}
}

type PreElement struct{ *Element }

func Pre(children ...ElementInterface) *PreElement {
	// Preformatted text
	return &PreElement{newElement("pre", children)}
}

type CommandElement struct{ *Element }

func Command(children ...ElementInterface) *CommandElement {
	// Command (HTML5)
	return &CommandElement{newElement("command", children)}
}

type ScriptElement struct{ *Element }

func Script(children ...ElementInterface) *ScriptElement {
	// Embedded script
	return &ScriptElement{newElement("script", children)}
}

type RtElement struct{ *Element }

func Rt(children ...ElementInterface) *RtElement {
	// Ruby text (HTML5)
	return &RtElement{newElement("rt", children)}
}

type SummaryElement struct{ *Element }

func Summary(children ...ElementInterface) *SummaryElement {
	// Summary, caption, or legend for a details control (HTML5)
	return &SummaryElement{newElement("summary", children)}
}
