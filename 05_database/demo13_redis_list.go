package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"reflect"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()
	_, err = conn.Do("LPUSH", "list1", "ele1", "ele2", "ele3", "ele4")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	//res, err := redis.String(conn.Do("LPOP", "list1"))//获取栈顶元素
	//res, err := redis.String(conn.Do("LINDEX", "list1", 3)) //获取指定位置的元素
	res, err := redis.Strings(conn.Do("LRANGE", "list1", 0, 3)) //获取指定下标范围的元素
	if err != nil {
		fmt.Println("redis POP error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %s \n", res)
	}
	//结果：
	//	connect success ...
	//	res type : string
	//  res  : [ele4 ele3 ele2 ele1]
}
