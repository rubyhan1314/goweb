package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	js, err :=simplejson.NewJson([]byte(`{
		"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
		}
		}`))
	if err != nil{
		fmt.Printf("error: %v\n", err)
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("%v\n",arr)
	fmt.Printf("%d\n",i)
	fmt.Printf("%s\n",ms)

}
