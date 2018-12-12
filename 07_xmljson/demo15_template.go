package main

import (
	"os"
	"html/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "ruby"}
	f2 := Friend{Fname: "王二狗"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
		an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
		my friend name is {{.Fname}}
		{{end}}
		{{end}}
		`)
	p := Person{UserName: "hanru",
		Emails: []string{"hanru@163.com", "hanru723@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
