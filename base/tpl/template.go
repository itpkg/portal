package tpl

import (
	"html/template"
	"io"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

type Model struct {
	Lang        string
	Url         string
	Author      string
	Keywords    string
	Description string
	Title       string
	Body        string
	BodyH       template.HTML
}

type Template func() (io.Writer, []*Model)

var templates = make([]Template, 0)

func Register(ts ...Template) {
	templates = append(templates, ts...)
}

func Dump(wrt io.Writer, view string, mod Model) error {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	mw := m.Writer("text/html", wrt)

	mod.BodyH = template.HTML(mod.Body)
	t, e := template.ParseFiles(view)
	if e != nil {
		return e
	}
	return t.Execute(mw, mod)
}
