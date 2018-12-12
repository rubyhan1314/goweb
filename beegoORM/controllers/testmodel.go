package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ModelController struct {
	beego.Controller
}

//-----------定义struct-------------

type User struct {
	// 对应user表
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	//注册驱动：如果是默认的三个可以不写
	orm.RegisterDriver("mysql", orm.DRMySQL) //可以不加

	//注册默认数据库，ORM 必须注册一个别名为default的数据库，作为默认使用。
	/*
	参数一：数据库别名
	参数二：驱动名称
	参数三：数据库连接字符串:username:password@tcp(127.0.0.1:3306)/databasename?charset=utf8
	参数四：设置数据库的最大空闲连接
	*/
	orm.RegisterDataBase("default", "mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8", 30)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag), new(UserInfo))
	orm.RegisterModel()
}

func (c *ModelController) CreateTable() {
	//自动建表
	orm.RunSyncdb("default", false, true)
	datainit()
}

func datainit() {
	o := orm.NewOrm()
	//rel  : 自动生成外键为 表名_id
	sql1 := "insert into user (name,profile_id) values ('hanru',1),('ruby',2),('王二狗',3);"
	sql2 := "insert into profile (age) values (20),(19),(21);"
	sql3 := "insert into tag (name) values ('offical'),('beta'),('dev');"
	sql4 := "insert into post (title,user_id) values ('paper1',1),('paper2',1),('paper3',2),('paper4',3),('paper5',3);"
	// m2m 生成的 表名：子表_主表s  主键自增
	sql5 := "insert into post_tags (tag_id, post_id) values (1,1),(1,3),(2,2),(3,3),(2,4),(3,4),(3,5); "

	//使用Raw（）.Exec（）执行sql
	o.Raw(sql1).Exec()
	o.Raw(sql2).Exec()
	o.Raw(sql3).Exec()
	o.Raw(sql4).Exec()
	o.Raw(sql5).Exec()
}

//------------------------简单的CRUD-----------------------------

/**
由于model定义为UserInfo，那么实际上操作的表示：user_info
 */
type UserInfo struct {
	Id       int64  `orm:"column(id)"`       // 也可以省略不写，orm会自动映射
	Username string `orm:"column(username)"` // 也可以省略不写，orm会自动映射
	Password string
}

func (c *ModelController) Get() {
	o := orm.NewOrm()
	o.Using("default") // 可以省略不写。你可以使用Using函数指定其他数据库

	/**
	通过orm对象来进行数据库的操作，这种情况是必须要知道主键
	 */

	//user := UserInfo{Username:"",Password:"zhangsan123"}
	// 1. insert
	//id, err := o.Insert(&user) // 第一个返回值为自增健 Id 的值

	//2. update
	//user := UserInfo{Id: 4, Username: "lisi", Password: "lisi123"}
	//num, err := o.Update(&user) //第一个返回值为影响的行数
	//if err != nil {
	//处理err
	//}
	//c.Ctx.WriteString(fmt.Sprintf("num:%d", num))
	//3. delete

	//num, err := o.Delete(&UserInfo{Id: 4})

	// 4.read ,查询
	//user := UserInfo{Id:1}
	//err := o.Read(&user)
	//
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	c.Ctx.WriteString(fmt.Sprintf("id:%d,name:%s,password:%s\n", user.Id,user.Username,user.Password))
	//}

	// 5. readOrCreate()
	user := UserInfo{Username: "李小花", Password: "xiaohuazzzz"}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	created, id, err := o.ReadOrCreate(&user, "username", "password")
	if err != nil {
		// 处理err
	}

	c.Ctx.WriteString(fmt.Sprintf("created:%t, id:%d", created, id))

}

//高级查询
func (c *ModelController) Query() {
	orm.Debug = true //是否开启调试模式，调试模式下回打印出sql
	o := orm.NewOrm()
	o.Using("default") // 可以省略不写。你可以使用Using函数指定其他数据库

	// 获取 QuerySeter 对象，user_info 为表名
	//qs := o.QueryTable("user")

	// 也可以直接使用对象作为表名
	//user := new(User)
	//qs := o.QueryTable(user) // 返回 QuerySeter

	//1. 查询所有的数据
	/*
	//定义一个User类型的切片
	var users []*User
	num, err := qs.All(&users) // WHERE name = 'hanru'
	if err != nil {
		// 处理err
	}
	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Profile.Id) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")

	*/
	// 2.指定查询：

	//qs.Filter("name", "hanru") // WHERE name = 'hanru'
	//qs.Filter("name__exact", "hanru") // WHERE name = 'hanru'
	// 使用 = 匹配，大小写是否敏感取决于数据表使用的 collation
	//qs.Filter("profile_id", nil) // WHERE profile_id IS NULL
	/*
	var users []*User
	num, err := qs.Filter("profile__isnull", true).Filter("name", "hanru").All(&users)
	if err != nil {
		// 处理err
	}
	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Profile.Id) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")
*/

	//3. 自定义条件
	var maps []orm.Params
	num, err := o.QueryTable("user").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
		c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
		c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

		for _, m := range maps {
			c.Ctx.WriteString("<tr>" +
				"<td>" + fmt.Sprintf("%v", m["Id"]) + "</td>" +
				"<td>" + fmt.Sprintf("%v", m["Name"]) + "</td>" +
				"<td>" + fmt.Sprintf("%v", m["Profile"]) + "</td>" +
				"</tr>")

		}
		c.Ctx.WriteString("</table></html>")
	}

}

//原生查询
func (c *ModelController) QuerySQL() {
	orm.Debug = true //是否开启调试模式，调试模式下回打印出sql
	o := orm.NewOrm()
	o.Using("default") // 可以省略不写。你可以使用Using函数指定其他数据库

	// 1. 查询
	/*
	var maps [] orm.Params
	num,err := o.Raw("SELECT * FROM user_info").Values(&maps)

	if err != nil {
		//处理err
	}
	fmt.Printf("Result Nums: %d\n", num)
	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, m:= range maps {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", m["id"]) + "</td>" +
			"<td>" + fmt.Sprintf("%v", m["username"]) + "</td>" +
			"<td>" + fmt.Sprintf("%v", m["password"]) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")
	*/

	// 2. exec()
	/*
	res, err := o.Raw("UPDATE user_info SET username = ? where id = ?", "尼古拉斯","7").Exec()
	if err != nil {
		//处理err
		c.Ctx.WriteString( fmt.Sprintf("err:%v", err))
	}
	num, _ := res.RowsAffected()
	c.Ctx.WriteString( fmt.Sprintf("共影响了 num:%d 条数据。。", num))
	*/

	// 3. QueryRow
	/*
	var user UserInfo
	err := o.Raw("SELECT id, username, password FROM user_info WHERE id = ?", 1).QueryRow(&user)
	if err != nil{
		// 处理err
	}

	c.Ctx.WriteString( fmt.Sprintf("id:%d, username:%s, password:%s", user.Id,user.Username,user.Password))
	*/

	// 4. QueryRows
	/*
	var users []UserInfo
	num, err := o.Raw("SELECT id, username, password FROM user_info").QueryRows(&users)
	if err != nil{
		// 处理err
	}

	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Username) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Password) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")
	*/

	// 5. row to map
	/*
	res := make(orm.Params)
	num, err := o.Raw("SELECT username, password FROM user_info").RowsToMap(&res, "username", "password")
	if err != nil {

	}
	// res is a map[string]interface{}{
	//  "ruby": 123456,
	//  "韩茹": hanru123,
	//  ....
	// }
	c.Ctx.WriteString(fmt.Sprintf("<html>"+"共查询了 num:%d 条数据。。", num) + "<br/>")
	for k, v := range res {
		c.Ctx.WriteString(fmt.Sprintf("key:%v,value:%v", k, v) + "<br/>")

	}

	c.Ctx.WriteString("</html>")
	*/


}


// QueryBuilder查询
func (c *ModelController)QueryBuilder(){

	orm.Debug = true //是否开启调试模式，调试模式下回打印出sql
	var users []MyUser

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("user.id",
		"user.name",
		"profile.age").
		From("user").
		InnerJoin("profile").On("user.profile_id = profile.id").
		Where("age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	num,err := o.Raw(sql, 19).QueryRows(&users)
	if err != nil{
		// 处理err
	}

	c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
	c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
	c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

	for _, user := range users {
		c.Ctx.WriteString("<tr>" +
			"<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
			"<td>" + fmt.Sprintf("%v", user.Age) + "</td>" +
			"</tr>")

	}
	c.Ctx.WriteString("</table></html>")

}


type MyUser struct {
	Id int
	Name string
	Age int

}