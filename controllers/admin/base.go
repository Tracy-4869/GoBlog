package admin

import (
	"net/http"

	"github.com/astaxie/beego"
)

// 继承
type BaseController struct {
	beego.Controller
}

// 定义状态常量及密钥
const (
	KEY = "482d909db559cd0922da137487d6f7e4"
)

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
