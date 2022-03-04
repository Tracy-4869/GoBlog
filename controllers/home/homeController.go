package home

import (
	"fmt"
	"goblog/models"
	"log"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	articleList, errArticle := models.GetArticleList()
	if errArticle != nil {
		log.Printf("get article list error: %s", errArticle)
	}
	tagList, errTag := models.GetTagList()
	if errTag != nil {
		log.Printf("get tag list error: %s", errTag)
	}
	totalArticleCount, errTotal := models.GetTotalArticleCount()
	if errTotal != nil {
		log.Printf("get article count error: %s", errTotal)
	}
	fmt.Println(tagList)
	c.Data["tagList"] = tagList
	c.Data["total"] = totalArticleCount
	c.Data["articleList"] = articleList
	c.Layout = "home/index.html"
	c.TplName = "home/list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"] = "home/left.html"
}
