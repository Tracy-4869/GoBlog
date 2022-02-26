package admin

import (
	"goblog/models"
	"goblog/utils"
	"log"
)

type AdminController struct {
	BaseController
}

// 管理员列表
func (c *AdminController) List() {
	adminList, err := models.GetAdminList()
	if err != nil {
		log.Printf("get adminList error:%s", err)
	}
	c.Data["adminList"] = adminList
	c.Layout = "admin/admin/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"] = "admin/nav.html"
}

// 管理员添加
func (c *AdminController) Add() {
	if c.IsPost() {
		admin := models.Admin{}
		admin.PassWord = utils.Md5Text(admin.PassWord)
		if err := c.ParseForm(&admin); err != nil {
			log.Printf("parseform admin error: %s", err)
			c.RespJson(StatZero, "parseform admin error")
		}
		if err := models.AddAdminInfo(admin); err != nil {
			log.Printf("add admin error: %s", err)
			c.RespJson(StatZero, "add admin error")
		}
		c.RespJson(StatOne, "add admin successful")
	}

	c.Data["statusMap"] = StatusMap
	c.Data["adminLevelMap"] = adminLevelMap

	c.Layout = "admin/admin/add.html"
	c.TplName = "admin/header.html"
}
