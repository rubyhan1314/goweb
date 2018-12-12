package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type TestLoginController struct {
	beego.Controller
}

type UserInfo struct {
	Username string              //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
	Password string `form:"pwd"` //也可以指定form表单中对应的name值，如果不写将无法解析该参数的值
}

func (c *TestLoginController) Login() {
	//获取cookie
	username := c.Ctx.GetCookie("username")
	password := c.Ctx.GetCookie("password")





	//验证用户名和密码：
	if username != "" {
		c.Ctx.WriteString("Username:" + username + ",Password:" + password)
	} else {

		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:9527/testlogin" method="post">
									用户名：<input type ="text" name="Username" />
									<br/>
									密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
									<br/>
									<input type="submit" value="提交">

								</form></html>`)
	}
}

func (c *TestLoginController) Post() {
	u := UserInfo{}
	if err := c.ParseForm(&u); err != nil {
		log.Panic(err)
	}
	//设置cookie
	c.Ctx.SetCookie("username",u.Username,100,"/")
	c.Ctx.SetCookie("password",u.Password,100,"/")


	//设置session
	c.SetSession("username",u.Username)
	c.SetSession("password",u.Password)

	c.Ctx.WriteString("Username:" + u.Username + ",Password:" + u.Password)

}
