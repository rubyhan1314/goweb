package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	urlStr := "http://localhost:2000/login"
	client := http.Client{}
	param := url.Values{
		"username": {"hanru"},
	}
	//jsonStr, _ := json.Marshal(param)
	//requestBody := bytes.NewBufferString(string(jsonStr))

	requestBody := bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", urlStr, requestBody)
	if err != nil {
		log.Fatal(err)
	}

	/*
	cookie的添加
	方式一：
		request.Header.Set("Cookie", "name=hanru")

	方式二：
		request.AddCookie(Cookie)

	 */
	cookId := &http.Cookie{Name: "userId", Value: "1234"}
	cookName := &http.Cookie{Name: "name", Value: "kongyixueyuan"}
	request.AddCookie(cookId)
	request.AddCookie(cookName)
	//使用text/html就会出错
	//application/x-www-form-urlencoded
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
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
	fmt.Printf("request.Header:%+v\n", request.Header)
	fmt.Printf("request.Cookies:%+v\n", request.Cookies())
}
