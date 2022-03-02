package admin

import (
	"goblog/models"
	"log"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	articleList, err := models.GetArticleList()
	if err != nil {
		log.Printf("get article error: %s", err)
	}
	tagList, errTag := models.GetTagList()
	if errTag != nil {
		log.Printf("get tagList error: %s", errTag)
	}
	tagMap := make(map[uint]string)
	for _, v := range tagList {
		tagMap[v.Id] = v.Name
	}
	c.Data["articleList"] = articleList
	c.Data["statusMap"] = StatusMap
	c.Data["tagMap"] = tagMap

	c.Layout = "admin/article/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"] = "admin/nav.html"
}

func (c *ArticleController) Add() {
	
}
