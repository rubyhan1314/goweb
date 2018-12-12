package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string //未导出的字段，首字母是小写的
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person{UserName: "王二狗",email:"ergou@163.com"}
	t.Execute(os.Stdout, p)
}
