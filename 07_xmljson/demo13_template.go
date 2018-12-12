package main

import (
	"os"
	"html/template"
)

type Person struct {
	UserName string
}


func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "韩茹"}
	t.Execute(os.Stdout, p)
}