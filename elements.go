package elem

// ========== Document Structure ==========

func Body(props Attrs, children ...interface{}) *Element {
	return NewElement("body", props, children...)
}

func Head(props Attrs, children ...interface{}) *Element {
	return NewElement("head", props, children...)
}

func Html(props Attrs, children ...interface{}) *Element {
	return NewElement("html", props, children...)
}

func Title(props Attrs, children ...interface{}) *Element {
	return NewElement("title", props, children...)
}

// ========== Text Formatting and Structure ==========

func A(props Attrs, children ...interface{}) *Element {
	return NewElement("a", props, children...)
}

func Div(props Attrs, children ...interface{}) *Element {
	return NewElement("div", props, children...)
}

func H1(props Attrs, children ...interface{}) *Element {
	return NewElement("h1", props, children...)
}

func H2(props Attrs, children ...interface{}) *Element {
	return NewElement("h2", props, children...)
}

func H3(props Attrs, children ...interface{}) *Element {
	return NewElement("h3", props, children...)
}

func P(props Attrs, children ...interface{}) *Element {
	return NewElement("p", props, children...)
}

func Span(props Attrs, children ...interface{}) *Element {
	return NewElement("span", props, children...)
}

func Text(content string) string {
	return content
}

// ========== Lists ==========

func Li(props Attrs, children ...interface{}) *Element {
	return NewElement("li", props, children...)
}

func Ul(props Attrs, children ...interface{}) *Element {
	return NewElement("ul", props, children...)
}

// ========== Forms ==========

func Button(props Attrs, children ...interface{}) *Element {
	return NewElement("button", props, children...)
}

// ========== Hyperlinks and Multimedia ==========

func Img(props Attrs) *Element {
	return NewElement("img", props)
}

// ========== Meta Elements ==========

func Meta(props Attrs) *Element {
	return NewElement("meta", props)
}

func Script(props Attrs, children ...interface{}) *Element {
	return NewElement("script", props, children...)
}
