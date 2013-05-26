# Tingo
Tingo is a Go libary for generating HTML from Go code. I found `template/html` too simple (no way to extend templates, create blocks etc). If you've ever used Jinja2 or the Django template language, you should know what I'm talking about.

Creating a full template language is a lot of work, so I decided to try a different approach: Doing it directly from Go.

Tingo is far from done, but it can generate some simple HTML.

### Example
The example below is taken from one of the tests. It renders in 0.09ms on my computer, with auto escaping turned on.
```go
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
```

### License
Tingo is available under the MIT license.