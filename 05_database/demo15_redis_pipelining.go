package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	conn.Send("HSET", "user","name", "hanru","age","30")
	conn.Send("HSET", "user","sex","female")
	conn.Send("HGET", "user","age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n",res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n",res3)
	//结果：
	//connect success ...
	//Receive res1:0
	//Receive res2:0
	//Receive res3:30
}
