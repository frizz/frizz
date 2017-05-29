package site

// frizz
type Page struct {
	Title  string
	Body   string
	Height uint8
	Foo    *Bar
}

type Bar struct {
	Baz string
}
