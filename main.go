package main

import (
	"fmt"
	"html/template"
	"os"
)

type book struct {
	Title         string
	Author        string
	MaxPage       uint
	PublishedDate string
}

func main() {
	Books := []book{}
	sadBook := book{
		Title:   "인간 실격",
		Author:  "다자이 오사무",
		MaxPage: 191,
	}
	redBook := book{
		Title:         "살인자의 기억법",
		Author:        "김영하",
		MaxPage:       173,
		PublishedDate: "2013-03",
	}
	brownBook := book{
		Title:         "허상의 어릿광대",
		Author:        "히가시노 게이고",
		MaxPage:       554,
		PublishedDate: "2021-12",
	}
	Books = append(Books, sadBook, redBook, brownBook)
	const booktmpl = `
	Book List
	
	{{range .}}
	{{.Title}} written by {{.Author}} {{if .PublishedDate}} in {{.PublishedDate}}
	{{- end}}
	Maximum Page {{.MaxPage}}
	{{end}}
	{{template "footnote"}}
`
	tmpl, err := template.New("booktmpl").Parse(booktmpl)
	if err != nil {
		fmt.Println(err)
	}
	_, err = tmpl.Parse("{{define `footnote`}}Selected by Kyungsub Kim{{end}}")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(os.Stdout, Books)
}
