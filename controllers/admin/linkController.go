package admin

import (
	"goblog/models"
	"log"
	"net/http"
)

type LinkController struct {
	BaseController
}

func (c *LinkController) List() {
	linksList, err := models.GetLinkList()
	if err != nil {
		log.Printf("get linksList error:%s", err)
	}

	c.Data["linksList"] = linksList
	c.Data["statusMap"] = StatusMap
	c.Layout = "admin/links/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"] = "admin/nav.html"
}

func (c *LinkController) Add() {
	if c.IsPost() {
		link := models.Links{}
		if err := c.ParseForm(&link); err != nil {
			log.Printf("parse link error: %s", err)
			c.RespJson(StatZero, "添加失败")
		}
		if err := models.AddLinkInfo(link); err != nil {
			log.Printf("add link error: %s", err)
			c.RespJson(StatZero, "添加失败")
		}
		c.RespJson(StatOne, "添加成功")
	} else {
		c.Data["statusMap"] = StatusMap
		c.Layout = "admin/links/add.html"
		c.TplName = "admin/header.html"
	}
}

func (c *LinkController) Edit() {
	if c.IsPost() {
		id := c.GetString("id")
		name := c.GetString("name")
		url := c.GetString("url")
		status := c.GetString("status")
		if err := models.UpdateLinkById(id, name, url, status); err != nil {
			log.Printf("update link by id error: %s", err)
			c.RespJson(StatZero, "更新失败")
		}
		c.RespJson(StatOne, "编辑成功")
	} else {
		linkId := c.GetString("id")
		info, err := models.GetLinkInfoById(linkId)
		if err != nil {
			log.Printf("get link error: %s", err)
		}
		linkList, errLink := models.GetLinkList()
		if errLink != nil {
			log.Printf("get linkList error: %s", errLink)
		}

		c.Data["info"] = info
		c.Data["statusMap"] = StatusMap
		c.Data["linkList"] = linkList
		c.Layout = "admin/links/edit.html"
		c.TplName = "admin/header.html"
	}
}

func (c *LinkController) DeleteLink() {
	id, err := c.GetInt("id")
	if err != nil {
		log.Printf("get id error: %s", err)
	}
	if err := models.DeleteLinkById(id); err != nil {
		log.Printf("delete link error: %s", err)
	}
	c.Redirect("/links/list", http.StatusFound)
}
