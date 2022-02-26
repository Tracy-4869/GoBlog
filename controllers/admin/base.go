package admin

import (
	"net/http"

	"github.com/astaxie/beego"
)

// 继承
type BaseController struct {
	beego.Controller
}

type Resp struct {
	Stat uint8  `json:"stat"`
	Msg  string `json:"msg"`
}

// 定义状态常量及密钥
const (
	StatZero = 0
	StatOne  = 1
	KEY      = "482d909db559cd0922da137487d6f7e4"
)

// 状态定义
var StatusMap = map[uint8]string{
	1: "正常",
	2: "冻结",
}

// 管理员等级
var adminLevelMap = map[uint8]string{
	1: "超级管理员",
	2: "普通管理员",
} 

// 检查session
func (c *BaseController) Prepare() {
	sessionUserName := c.GetSession("userName")
	sessionToken := c.GetSession("token")
	if sessionUserName == "" || sessionToken == "" {
		c.Redirect("/login", http.StatusFound)
	}

	cookieUserName := c.Ctx.GetCookie("userName")
	cookieToken := c.Ctx.GetCookie("token")
	if cookieUserName == "" || cookieToken == "" {
		c.Redirect("/login", http.StatusFound)
	}

	if sessionUserName != cookieUserName || sessionToken != cookieToken {
		c.Redirect("/login", http.StatusFound)
	}
}

// 封装对接的json
func (c *BaseController) RespJson(stat uint8, msg string) {
	resp := Resp{
		Stat: stat,
		Msg:  msg,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// post判断
func (c *BaseController) IsPost() bool {
	return c.Ctx.Request.Method == "POST"
}
