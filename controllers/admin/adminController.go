package admin

import (
	"goblog/models"
	"log"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) List() {
	adminList, err := models.GetAdminList()
	if err != nil {
		log.Printf("get adminList error:%s", err)
	}
	c.Data["adminList"] = adminList
	c.Layout = "admin/admin/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"]    = "admin/nav.html"
}
