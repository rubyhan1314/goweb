package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
	连接数据库:func Open(driverName, dataSourceName string) (*DB, error)
	Open打开一个dirverName指定的数据库，dataSourceName指定数据源，
	一般包至少括数据库文件名和（可能的）连接信息。

	driverName: 使用的驱动名. 这个名字其实就是数据库驱动注册到 database/sql 时所使用的名字.
	dataSourceName: 数据库连接信息，这个连接包含了数据库的用户名, 密码, 数据库主机以及需要连接的数据库名等信息.

	drvierName,"mysql"
	dataSourceName,用户名:密码@协议(地址:端口)/数据库?参数=参数值

	 */

	//"root:hanru1314@tcp(127.0.0.1:3306)/ruby?charset=utf8"
	db, err := sql.Open("mysql", "root:hanru1314@tcp:(127.0.0.1:3306)/mytest?charset=utf8")
	fmt.Println(db)
	fmt.Println(err)
	if err != nil {
		fmt.Println("连接有误。。")
		return
	}
	fmt.Println("连接成功。。")
	db.Close()

}
