package main

//step1：导入包
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	/*
	查询操作：
	 */
	//step2:打开数据库，建立连接
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")

	//stpt3：查询数据库
	rows, err := db.Query("SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("查询有误。。")
		return
	}

	//fmt.Println(rows.Columns()) //[uid username departname created]
	//定义一个map，用于存储从数据库中查询出来的数据，字段作为key，string，数据作为value，任意类型，空接口
	map1 := make(map[string]interface{})
	datas := make([] map[string]interface{}, 0)

	//step4：操作结果集获取数据
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		if err := rows.Scan(&uid, &username, &departname, &created); err != nil {
			fmt.Println("获取失败。。")
		}
		//将读取到的数据，存入了map中
		map1["uid"] = uid
		map1["username"] = username
		map1["departname"] = departname
		map1["created"] = created
		//将map存入切片中
		datas = append(datas, map1)

	}
	//step5：关闭资源
	rows.Close()
	db.Close()


	for _, v := range datas {
		fmt.Println(v)
	}

}

/*
查询：处理查询后的结果：
	思路一：将数据，存入到map中
	思路二：创建结构体：
 */
