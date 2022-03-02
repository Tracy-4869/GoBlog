package admin

import (
	"goblog/models"
	"log"
	"net/http"
)

type TagController struct {
	BaseController
}

func (c *TagController) List() {
	tagList, err := models.GetTagList()
	if err != nil {
		log.Printf("get tagList error:%s", err)
	}
	c.Data["tagList"] = tagList
	c.Data["statusMap"] = StatusMap
	c.Layout = "admin/tag/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"] = "admin/nav.html"
}

func (c *TagController) Add() {
	if c.IsPost() {
		tag := models.ArticleTag{}
		if err := c.ParseForm(&tag); err != nil {
			log.Printf("parse form error: %s", err)
		}
		if err := models.AddTag(tag); err != nil {
			log.Printf("add tag error: %s", err)
			c.RespJson(StatZero, "add tag error")
		}
		c.RespJson(StatOne, "add tag successful")
	}
	c.Data["statusMap"] = StatusMap
	c.Layout = "admin/tag/add.html"
	c.TplName = "admin/header.html"
}

func (c *TagController) Edit() {
	if c.IsPost() {
		id     := c.GetString("id")
		name   := c.GetString("name")
		status := c.GetString("status")

		if err := models.UpdateTagById(id, name, status); err != nil {
			log.Printf("update tag error: %s", err)
		}
		c.RespJson(StatOne, "编辑成功")
	} else {
		id := c.GetString("id")
		tagInfo, err := models.GetTagInfoById(id)
		if err != nil {
			log.Printf("get tag info error: %s", err)
		}
		c.Data["tagInfo"] = tagInfo
		c.Data["statusMap"] = StatusMap
		c.Layout = "admin/tag/edit.html"
		c.TplName = "admin/header.html"
	}
}

func (c *TagController)DeleteTag() {
	id, err := c.GetInt("id")
	if err != nil {
		log.Printf("get id error: %s", err)
	}
	if err := models.DeleteTagById(id); err != nil {
		log.Printf("delete tag error: %s", err)
	}
	c.Redirect("/tag/list", http.StatusFound)
}
