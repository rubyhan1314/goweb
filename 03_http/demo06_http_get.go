package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
