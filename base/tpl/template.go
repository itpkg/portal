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
	SubTitle    string
	Body        template.HTML
}

func (p *Model) SetBody(body string) {
	p.Body = template.HTML(body)
}

func Dump(wrt io.Writer, view string, mod *Model) error {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	mw := m.Writer("text/html", wrt)
	defer mw.Close()

	t, e := template.ParseFiles(view)
	if e != nil {
		return e
	}
	return t.Execute(mw, mod)

}

func DumpF(wrt io.Writer, view string, mod *Model) error {
	t, e := template.ParseFiles(view)
	if e != nil {
		return e
	}
	return t.Execute(wrt, mod)

}
