package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, err := http.Post("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		"application/x-www-form-urlencoded",
		strings.NewReader("theCityName=苏州"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
