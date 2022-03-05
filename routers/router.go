package routers

import (
	"goblog/controllers/admin"
	"goblog/controllers/home"

	"github.com/astaxie/beego"
)

func Init() {
	// 后端首页
	beego.Router("/login", &admin.IndexController{}, "get:Index")
	// 登录操作
	beego.Router("/login", &admin.IndexController{}, "post:Login")
	// 登录成功主页
	beego.Router("/main", &admin.MainController{}, "get:Index")
	// 管理员
	beego.Router("/admin/list", &admin.AdminController{}, "get:List")
	beego.Router("/admin/add", &admin.AdminController{}, "get,post:Add")
	beego.Router("/admin/edit", &admin.AdminController{}, "get,post:Edit")
	beego.Router("/admin/delete", &admin.AdminController{}, "get,delete:DeleteAdmin")
	// 标签
	beego.Router("/tag/list", &admin.TagController{}, "get:List")
	beego.Router("/tag/add", &admin.TagController{}, "get,post:Add")
	beego.Router("/tag/edit", &admin.TagController{}, "get,post:Edit")
	beego.Router("/tag/delete", &admin.TagController{}, "get:DeleteTag")
	// 文章
	beego.Router("/article/list", &admin.ArticleController{}, "get:List")
	beego.Router("/article/add", &admin.ArticleController{}, "get,post:Add")
	beego.Router("/article/edit", &admin.ArticleController{}, "get,post:Edit")
	beego.Router("/article/delete", &admin.ArticleController{}, "get:DeleteArticle")
	// 个人简介

	// 链接

	// 前台首页
	beego.Router("/", &home.HomeController{}, "get:Index")

}
