package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/setcookie", SetCookie)
	http.HandleFunc("/getcookie", GetCookie)
	http.ListenAndServe("localhost:8080", nil)

}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	//设置cookie
	cookie := http.Cookie{Name: "name", Value: "hanru", Path: "/", MaxAge: 60}
	http.SetCookie(w, &cookie)
	w.Write([]byte("write cookie ok"))

}

//Go读取cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie2, _ := r.Cookie("name")
	fmt.Fprint(w, cookie2)
	//还有另外一种读取方式
	//for _, cookie := range r.Cookies() {
	//	fmt.Fprint(w, cookie.Name)
	//}
}
