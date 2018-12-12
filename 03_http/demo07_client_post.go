package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/url"
)

func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName"
	client := http.Client{}
	param := &url.Values{
		"theCityName": {"苏州"},
	}
	requestBody := bytes.NewBufferString(param.Encode())
	response, err := client.Post(urlStr, "application/x-www-form-urlencoded", requestBody)
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
	fmt.Printf("%+v", response)
}