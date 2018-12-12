package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)


func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()
}