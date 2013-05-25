package tingo

import (
	"fmt"
	"testing"
)

func ExampleBody() {
	fmt.Println(Body().Render())
	// Output: <body></body>
}

func BenchmarkBody(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body().Render()
	}
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body(
			Header(
				H1().Prepend("Hello World!"),
			).Id("top"),
			Section(
				P().Class("message").Prepend("This is a benchmark."),
			).Id("main"),
			Footer(
				Span().Class("credits").Prepend("Generated with Tingo"),
			).Id("bottom"),
		).Render()
	}
}

func BenchmarkBasicIndented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Body(
			Header(
				H1().Prepend("Hello World!"),
			).Id("top"),
			Section(
				P().Class("message").Prepend("This is a benchmark."),
			).Id("main"),
			Footer(
				Span().Class("credits").Prepend("Generated with Tingo"),
			).Id("bottom"),
		).RenderIndent("\t")
	}
}

func TestRender(t *testing.T) {
	Body(
		Div(
			P().Class("green"),
		).Id("main").Prepend("Lorem ipsum...").Append("Sit amet"),
	).Render()
}

func TestRenderIndent(t *testing.T) {
	Body(
		Div(
			P().Class("green"),
		).Id("main").Prepend("Lorem ipsum...").Append("Sit amet"),
	).RenderIndent("  ")

	fmt.Println(Html(
		Head(
			Title().Prepend("Test"),
		),
		Body(
			H1().Prepend("This is a test"),
			Input().Type("text"),
		),
	).RenderIndent("\t"))
}
