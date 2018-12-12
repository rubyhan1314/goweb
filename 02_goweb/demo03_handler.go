package main

import (
	"net/http"
	"io"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hanru", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "你好，韩茹。。")
	})
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "byebye")
	})

	mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://www.baidu.com", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/hello", sayhello)
	http.ListenAndServe(":8080", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}