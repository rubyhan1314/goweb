package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")

	fmt.Println("username:", template.HTMLEscapeString(username)) //输出到服务器端

	fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

	template.HTMLEscape(w, []byte(username)) //输出到客户端

	//fmt.Fprintf(w, username) //这个写入到w的是输出到客户端的




}

func login2(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Println(username)

	//进行模板解析
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(w, "T", username)
	err = t.ExecuteTemplate(w, "T", template.HTML(username))


	//如果转义失败 抛出对应错误 终止程序
	if err != nil {
		log.Fatal(err)
	}
}

func main() {


	http.HandleFunc("/login", login2)      //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
