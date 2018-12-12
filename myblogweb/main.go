package main

import (
	_ "myblogweb/routers"
	"github.com/astaxie/beego"
	"myblogweb/utils"
)

func main() {
	utils.InitMysql()
	//beego.BConfig.WebConfig.Session.SessionOn = true // 打开session
	beego.Run()
}

