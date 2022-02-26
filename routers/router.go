package routers

import (
	"goblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	// 后端首页
	beego.Router("/login", &admin.IndexController{}, "get:Index")
	// 登录操作
	beego.Router("/login", &admin.IndexController{}, "post:Login")
	// 登录成功主页
	beego.Router("/main", &admin.MainController{}, "get:Index")
	// 管理员列表
	beego.Router("/admin/list", &admin.AdminController{}, "get:List")
	// 管理员添加
	beego.Router("/admin/add", &admin.AdminController{}, "get,post:Add")

}
