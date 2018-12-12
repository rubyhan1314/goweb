package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"bytes"
	"log"
)

func main() {
	/**
	也可以使用自己搭建的服务器
	 */
	urlStr := "http://localhost:2001/login"
	param := url.Values{
		"username": {"孔壹学院"},
		"password": {"hanru12345"},
	}
	requestBody := bytes.NewBufferString(param.Encode())
	response, err := http.Post(urlStr, "application/x-www-form-urlencoded", requestBody)

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
