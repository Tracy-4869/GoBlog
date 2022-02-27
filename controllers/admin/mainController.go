package admin

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.Data["userName"] = c.GetSession("userName")
	c.Layout = "admin/main.html"
	c.TplName = "admin/header.html"
}
