package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
)

func main() {
	/**
	还是使用上一个服务器
	使用http.PostForm()
	 */
	response, err := http.PostForm("http://localhost:2001/login",
		url.Values{"username": {"孔壹学院"}, "password": {"kongyixueyuan"}})

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("response:%+v\n", response)
	fmt.Printf("response.Header:%+v\n", response.Header)
	fmt.Printf("response.Cookies:%+v\n", response.Cookies())
}
