package main

import (
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	// Must 它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo:{{if ``}}不回输出.{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo:{{if `anything`}} if 部分 {{else}} else 部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
