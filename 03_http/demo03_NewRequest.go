package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	/**
	客户端实现方式一：
	1.先生成http.client ->
	2.再生成 http.request ->
	3.之后提交请求：client.Do(request) ->
	4.处理返回结果，每一步的过程都可以设置一些具体的参数。
	 */
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"
	//step1：创建client对象
	client := http.Client{}
	//step2：创建request
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Fatal(err)
	}
	//step3：发送请求
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

	fmt.Println("-----------------------------------------")
	fmt.Printf("response：%+v\n", response)
	fmt.Println("-----------------------------------------")
	fmt.Printf("response.Body：%+v\n", response.Body)
	fmt.Printf("response.Header：%+v\n", response.Header)
	fmt.Printf("response.StatusCode：%+v\n", response.StatusCode)
	fmt.Printf("response.Status：%+v\n", response.Status)
	fmt.Printf("response.Request：%+v\n", response.Request)
	fmt.Printf("response.Cookies：%+v\n", response.Cookies())
}