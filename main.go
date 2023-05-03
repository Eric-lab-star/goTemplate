package main

import (
	"os"
	"text/template"
)

type Invetory struct {
	Material string
	Count    uint
}

func (i *Invetory) Shout(name string) string {
	return name + "is made by master kim"
}
func main() {
	lists := []int{1, 2, 3, 4}
	tmpl, err := template.New("trivial").Parse("{{range $elem := .}} {{$elem}} {{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, lists)
	if err != nil {
		panic(err)
	}
}
