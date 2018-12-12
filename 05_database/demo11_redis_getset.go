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
	_, err = conn.Do("SET", "name", "hanru")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}
}