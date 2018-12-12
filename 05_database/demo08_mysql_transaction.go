package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
	事务：
		4大特性：ACID
		原子性：
		一致性：
		隔离性：
		永久性：
	 */
	//ruby-->王二狗,2000元
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	//开启事务
	tx, _ := db.Begin()
	//提供一组sql操作
	var aff1, aff2 int64 = 0, 0
	result1, _ := tx.Exec("UPDATE account SET money=1000 WHERE id=?", 1)
	result2, _ := tx.Exec("UPDATE account SET money=3000 WHERE id=?", 2)
	//fmt.Println(result2)
	if result1 != nil {
		aff1, _ = result1.RowsAffected()
	}
	if result2 != nil {
		aff2, _ = result2.RowsAffected()
	}
	fmt.Println(aff1)
	fmt.Println(aff2)

	if aff1 == 1 && aff2 == 1 {
		//提交事务
		tx.Commit()
		fmt.Println("操作成功。。")
	} else {
		//回滚
		tx.Rollback()
		fmt.Println("操作失败。。。回滚。。")
	}

}
