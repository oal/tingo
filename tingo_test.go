package tingo

import (
	"fmt"
	"testing"
)

func BenchmarkBody(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body().Render()
	}
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body(
			Header(
				H1().TextPrepend("Hello World!"),
			).Id("top"),
			Section(
				P().Class("message").TextPrepend("This is a benchmark."),
			).Id("main"),
			Footer(
				Span().Class("credits").TextPrepend("Generated with Tingo"),
			).Id("bottom"),
		).Render()
	}
}

func BenchmarkBasicIndented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body(
			Header(
				H1().TextPrepend("Hello World!"),
			).Id("top"),
			Section(
				P().Class("message").TextPrepend("This is a benchmark."),
			).Id("main"),
			Footer(
				Span().Class("credits").TextPrepend("Generated with Tingo"),
			).Id("bottom"),
		).RenderIndent("\t")
	}
}

func BenchmarkBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Html(
			Head(
				Title().TextAppend("Page title"),
				Style().TextAppend("body { font-family: Arial, sans-serif }"),
			),
			Body(
				Header(
					H1(A().TextPrepend("Page title")),
					Nav(
						Ul(
							Li(A().TextPrepend("Home")),
							Li(A().TextPrepend("About")),
							Li(A().TextPrepend("Services")),
						),
					),
				),
				Section(
					H2().TextPrepend("Lorem ipsum"),
					P().TextPrepend("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent cursus odio sit amet metus blandit porttitor. Etiam convallis nibh eget erat cursus molestie auctor purus semper. Ut at suscipit mauris. Duis interdum nunc sodales neque consectetur fringilla. In hac habitasse platea dictumst. Integer nec ante purus. Morbi eu arcu sit amet felis fermentum adipiscing. Proin imperdiet lorem sed diam viverra at vestibulum libero imperdiet. Duis tortor orci, interdum eu eleifend blandit, dapibus non risus. Phasellus aliquam risus quis orci porta eleifend blandit diam mollis. Vivamus eu dolor mi, nec consequat diam. Pellentesque vel neque metus. Phasellus ac nisi orci, dapibus posuere diam. Nulla adipiscing, arcu in tempus facilisis, purus tellus bibendum massa, at consectetur odio lectus a lacus."),
				).Id("main"),
				Footer(
					P().TextPrepend("Tingo test page.").Class("credits"),
				),
			),
		).RenderIndent("\t")
	}
}

func TestRender(t *testing.T) {
	Body(
		Div(
			P().Class("green"),
		).Id("main").TextPrepend("Lorem ipsum...").TextAppend("Sit amet"),
	).Render()
}

func TestRenderIndent(t *testing.T) {
	Body(
		Div(
			P().Class("green"),
		).Id("main").TextPrepend("Lorem ipsum...").TextAppend("Sit amet"),
	).RenderIndent("  ")

	fmt.Println(Html(
		Head(
			Title().TextPrepend("Test"),
		),
		Body(
			H1().TextPrepend("This is a test <strong>escaped</strong>"),
			TextInput("name"),
		),
	).RenderIndent("\t"))
}
