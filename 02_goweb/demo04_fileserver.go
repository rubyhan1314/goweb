package main

import (
	"log"
	"net/http"
)

func main() {
	/**
	FileServer：
　　
	1.www.xx.com/ 根路径 直接使用

　　　　　http.Handle("/", http.FileServer(http.Dir("/tmp")))

　　
　　2.www.xx.com/c/ 带有请求路径的 需要添加函数
　　　　　　http.Handle("/c/", http.StripPrefix("/c/", http.FileServer(http.Dir("/tmp"))))
	 */
	err := http.ListenAndServe(":2003", http.FileServer(http.Dir("/Users/ruby/Documents/ChainDesk/GoWeb开发实战")))
	if err != nil {
		log.Fatal(err)
	}
}
