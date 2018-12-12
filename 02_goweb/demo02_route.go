package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

/**
ServeHTTP方法，
他需要2个参数，一个是http.ResponseWriter，另一个是http.Request
往http.ResponseWriter写入什么内容，浏览器的网页源码就是什么内容。
http.Request里面是封装了，浏览器发过来的请求（包含路径、浏览器类型等等）。
 */
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my route!")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)
}
