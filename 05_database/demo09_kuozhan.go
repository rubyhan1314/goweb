package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var Db *sqlx.DB

type User struct {
	Uid        int
	Username   string
	Departname string
	Created    string
}

func main() {

	db, err := sqlx.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = db
	var users []User
	err = Db.Select(&users, "SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("this is Select res:%v\n", users)
	var user User
	err1 := Db.Get(&user, "SELECT uid,username,departname,created FROM userinfo where uid = ?", 1)
	if err1 != nil {
		fmt.Println("GET error :", err1)
	} else {
		fmt.Printf("this is GET res:%v", user)
	}
}

//this is Select res:[{1 韩茹 技术部 2018-11-21} {2 王二狗 行政部 2018-11-11}]
//this is GET res:{1 韩茹 技术部 2018-11-21}
