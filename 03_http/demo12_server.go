package main

import (
	"log"
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("body", string(body))
	fmt.Printf("r:%+v\n", r)
	fmt.Printf("Request：Header：%+v\n", r.Header)
	fmt.Printf("cookies:%+v\n", r.Cookies())
	cookid, err := r.Cookie("userId")
	if err == nil {
		fmt.Println(cookid.Name, cookid.Value)
	}

	if len(r.Form["username"]) > 0 {
		username := r.Form["username"][0]
		fmt.Println("username:", username)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	io.WriteString(w, "login success ....")
}

func main() {
	http.HandleFunc("/login", HandleRequest)
	err := http.ListenAndServe(":2000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
