# Tingo
Tingo is a Go libary for generating HTML from Go code. I found `template/html` too simple (no way to extend templates, create blocks etc). If you've ever used Jinja2 or the Django template language, you should know what I'm talking about.

Creating a full template language is a lot of work, so I decided to try a different approach: Doing it directly from Go.

Tingo is far from done, but it can generate some simple HTML.

### Example
    debug := true
    Html(
		Head(
			Title().Before("Hello World"),
		),
		Body(
			H1().Before("Hello World"),
			P().Before("This is a test document.").If(debug), // Will only be shown if debug is true.
		),
	).RenderIndent("\t")

### License
Tingo is available under the MIT license.