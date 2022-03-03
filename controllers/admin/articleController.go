package admin

import (
	"goblog/models"
	"log"
	"net/http"
	"strconv"
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
	if c.IsPost() {
		article := models.Article{}
		if err := c.ParseForm(&article); err != nil {
			log.Printf("parse form error: %s", err)
			c.RespJson(StatZero, "添加失败")
		}
		if err := models.AddArticle(article); err != nil {
			log.Printf("add article error: %s", err)
			c.RespJson(StatZero, "添加失败")
		}
		// 文章状态正常，更新该标签下文章数量
		if article.Status == 1 {
			if err := models.AddTotalById(article.TagId); err != nil {
				log.Printf("add tag total error: %s", err)
			}
		}
		c.RespJson(StatOne, "添加成功")
	} else {
		tagList, err := models.GetTagList()
		if err != nil {
			log.Printf("get tag list error: %s", err)
		}
		c.Data["tagList"] = tagList
		c.Data["statusMap"] = StatusMap
		c.Layout = "admin/article/add.html"
		c.TplName = "admin/header.html"
	}
}

func (c *ArticleController) Edit() {
	articleId := c.GetString("id")
	article, err := models.GetArticleInfoById(articleId)
	if err != nil {
		log.Printf("get article error: %s", err)
	}
	if c.IsPost() {
		id := c.GetString("id")
		title := c.GetString("title")
		content := c.GetString("content")
		author := c.GetString("author")
		status := c.GetString("status")
		tag_id := c.GetString("tag_id")
		if err := models.UpdateArticleById(id, tag_id, title, content, author, status); err != nil {
			log.Printf("update article by id error: %s", err)
			c.RespJson(StatZero, "更新失败")
		}
		// 对应标签数量更改
		tagId := strconv.Itoa(int(article.TagId))
		if tagId != tag_id {
			// 状态
			if article.Status == 1 {
				if errTagId := models.MinusTotalById(tagId); errTagId != nil {
					log.Printf("minus total error: %s", errTagId)
				}
			}
			if status == "1" {
				if errTagNewId := models.AddTotalByStringId(tag_id); errTagNewId != nil {
					log.Printf("add total error: %S", errTagNewId)
				}
			}
		} else {
			if status == "1" && article.Status == 2 {
				if errTagNewId := models.AddTotalByStringId(tag_id); errTagNewId != nil {
					log.Printf("add total error: %S", errTagNewId)
				}
			} else if status == "2" && article.Status == 1 {
				if errTagId := models.MinusTotalById(tagId); errTagId != nil {
					log.Printf("minus total error: %s", errTagId)
				}
			}
		}
		c.RespJson(StatOne, "更新成功")
	} else {
		tagList, errTag := models.GetTagList()
		if errTag != nil {
			log.Printf("get tagList error: %s", errTag)
		}

		c.Data["article"] = article
		c.Data["statusMap"] = StatusMap
		c.Data["tagList"] = tagList
		c.Layout = "admin/article/edit.html"
		c.TplName = "admin/header.html"
	}
}

func (c *ArticleController) DeleteArticle() {
	id, err := c.GetInt("id")
	if err != nil {
		log.Printf("get id error: %s", err)
	}
	// 标签数量更改
	strId := strconv.Itoa(id)
	if article, err := models.GetArticleInfoById(strId); err != nil {
		log.Printf("get article by id error: %s", err)
	} else {
		// 正常状态减1, 冻结状态不用操作
		if article.Status == 1 {
			if errMinus := models.MinusTotalById(strconv.Itoa(int(article.TagId))); errMinus != nil {
				log.Printf("minus total error: %s", errMinus)
			}
		}
	}
	if err := models.DeleteArticleById(uint(id)); err != nil {
		log.Printf("delete article error: %s", err)
	}
	c.Redirect("/article/list", http.StatusFound)
}
