package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	/**
	方式二 调用client的API

	1.先生成client，
	2.之后用client.get/post..

	还额外添加了req.Header.Set("Content-Type", bodyType)...
	 */
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"
	//step1：创建client对象
	client := http.Client{}

	//step2：调用get方法，发送请求
	response, err := client.Get(urlStr)
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