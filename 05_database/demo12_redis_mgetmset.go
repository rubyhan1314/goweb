package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"reflect"
)


func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()
	_, err = conn.Do("MSET", "name", "hanru","age",30)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Strings(conn.Do("MGET", "name","age"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("MGET name: %s \n", res)
		fmt.Println(len(res))
	}
	//结果：
	//connect success ...
	//res type : []string
	//MGET name: [hanru 30]
	//2
}