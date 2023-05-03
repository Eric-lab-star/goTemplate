package main

import (
	"os"
	"text/template"
)

type Invetory struct {
	Material string
	count    uint
}

func main() {
	sweater := Invetory{"wool", 12}
	tmpl, err := template.New("trivial").Parse("{{.count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweater)
}
