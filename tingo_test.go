package tingo

import (
	"fmt"
	"testing"
)

func ExampleBody() {
	fmt.Println(Body().Render())
	// Output: <body></body>
}

func TestRender(t *testing.T) {
	Body(
		Div(
			//Ul().Loop(Li, []string{"First list item", "Second one", "And the third!"}),
			P().Class("green"),
		).Id("main").Before("Lorem ipsum...").After("Sit amet"),
	).Render()
}

func TestRenderIndent(t *testing.T) {
	Body(
		Div(
			//Ul().Loop(Li, []string{"First list item", "Second one", "And the third!"}),
			P().Class("green"),
		).Id("main").Before("Lorem ipsum...").After("Sit amet"),
	).RenderIndent("  ")

	fmt.Println(Html(
		Head(
			Title().Before("Test"),
		),
		Body(
			H1().Before("This is a test"),
			Input().Type("text"),
		),
	).RenderIndent("\t"))
}
