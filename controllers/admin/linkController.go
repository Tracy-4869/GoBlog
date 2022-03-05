package admin

type LinkController struct {
	BaseController
}

func (c *LinkController) List() {

	c.Layout  = "admin/links/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"]    = "admin/nav.html"
	c.LayoutSections["footer"] = "admin/footer.html"
}

func (c *LinkController) Add() {
	c.Layout   = "admin/links/add.html"
	c.TplName  = "admin/header.html"
}

func (c *LinkController) Edit() {
	c.Layout   = "admin/links/edit.html"
	c.TplName  = "admin/header.html"
}
