package main
// step1：导入包
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

func main()  {
	// step2：打开数据库，相当于和数据库建立连接：db对象
	/*
	func Open(driverName, dataSourceName string) (*DB, error)
	drvierName,"mysql"
	dataSourceName,用户名:密码@协议(地址:端口)/数据库?参数=参数值
	 */
	db,err:=sql.Open("mysql","root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	if err !=nil{
		fmt.Println("连接失败。。")
		return
	}

	//step3：修改一条数据，我们直接使用Exec()方法。
	//更新数据

	result, err := db.Exec("UPDATE userinfo SET username = ?, departname = ? WHERE uid = ?", "王二狗","行政部",2)
	if err !=nil{
		fmt.Println("更新数据失败。。",err)
	}

	//step4：处理sql操作后的结果
	lastInsertId,err:=result.LastInsertId()
	rowsAffected,err:=result.RowsAffected()
	fmt.Println("lastInsertId",lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)


	//step5：关闭资源
	db.Close()


}
