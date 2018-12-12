package routers

import (
	"beegoORM/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/testcreatetable", &controllers.ModelController{},"get:CreateTable")
    beego.Router("/testmodel", &controllers.ModelController{},"get:Get;post:Post")
    beego.Router("/testquery", &controllers.ModelController{},"get:Query;post:Post")
    beego.Router("/testquerysql", &controllers.ModelController{},"get:QuerySQL;post:Post")
    beego.Router("/testquerybuilder", &controllers.ModelController{},"get:QueryBuilder;post:Post")
}
