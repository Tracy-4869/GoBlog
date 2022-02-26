package admin

import (
	"goblog/models"
	"goblog/utils"
	"log"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type Resp struct {
	Stat uint8  `json:"stat"`
	Msg  string `json:"msg"`
}

// 后台首页 登录页面
func (c *IndexController) Index() {
	c.Layout = "admin/index.html"
	c.TplName = "admin/header.html"
}

// 登录
func (c *IndexController) Login() {
	resp := Resp{
		Stat: 0,
		Msg:  "用户名或密码错误",
	}
	// 获取用户名和密码
	userName := c.GetString("username")
	passWord := c.GetString("password")

	// 通过用户名获取用户信息
	admin, err := models.GetAdminInfoByName(userName)
	if err != nil {
		log.Printf("db error: %s", err)
	}

	// 密码验证
	md5Pass := utils.Md5Text(passWord)
	if admin.PassWord == md5Pass {
		resp.Msg = "登录成功"
		resp.Stat = 1
	}

	// 保存session和cookie
	token := c.createToken(userName, passWord)
	c.setLoginSession(userName, token)
	c.setLoginCookie(userName, token)
	
	c.Data["json"] = resp
	c.ServeJSON()
}

// 设置token
func (c *IndexController) createToken(userName, passWord string) string {
	// 拼接密钥
	str := KEY + userName + passWord
	return utils.Md5Text(str)
}

// 设置session
func (c *IndexController) setLoginSession(userName, token string) {
	c.SetSession("userName", userName)
	c.SetSession("token", token)
}

// 设置cookie
func (c *IndexController) setLoginCookie(userName, token string) {
	c.Ctx.SetCookie("userName", userName)
	c.Ctx.SetCookie("token", token)
}