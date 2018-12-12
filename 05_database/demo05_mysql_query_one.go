package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	/*
	查询一条
	 */
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")

	row := db.QueryRow("SELECT uid,username,departname,created FROM userinfo WHERE uid=?", 1)
	var uid int
	var username, departname, created string
	/*
	row：Scan()-->将查询的结果从row取出
		err对象
		判断err是否为空，
			为空，查询有结果，数据可以成功取出
			不为空，没有数据，sql: no rows in result set
	 */
	err := row.Scan(&uid, &username, &departname, &created)

	//fmt.Println(err)
	if err != nil {
		fmt.Println("查无数据。。")
	} else {
		fmt.Println(uid, username, departname, created)

	}
	db.Close()

}
