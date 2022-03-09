package home

import (
	"goblog/models"
	"log"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	articleList := getArticleList()
	for k, v := range articleList {
		absRune := []rune(v.Content)
		articleList[k].Abstract = string(absRune[:16])
	}

	c.Data["tagList"] = getTagList()
	c.Data["total"] = getTotalArticleCount()
	c.Data["articleList"] = articleList
	c.Data["links"] = getLinkList()
	c.Data["profile"] = getProfile()
	c.Data["topArticle"] = getTopArticle()
	c.Data["topTenViewArt"] = getArticleTopThree()

	c.Layout = "home/index.html"
	c.TplName = "home/list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"] = "home/left.html"
}

func (c *HomeController) Detail() {
	id := c.Ctx.Input.Param(":id")

	c.Data["total"] = getTotalArticleCount()
	c.Data["tagList"] = getTagList()
	c.Data["links"] = getLinkList()
	c.Data["topArticle"] = getTopArticle()
	c.Data["profile"] = getProfile()
	c.Data["topTenViewArt"] = getArticleTopThree()
	c.Data["articleContent"] = getArticleInfoById(id)

	// 更新文章阅读量
	updateArticleReadNum(id)

	c.Layout = "home/detail.html"
	c.TplName = "home/content.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"] = "home/left.html"
}

func (c *HomeController) Category() {
	tagId := c.Ctx.Input.Param(":id")

	c.Data["tagId"] = tagId
	c.Data["links"] = getLinkList()
	c.Data["topArticle"] = getTopArticle()
	c.Data["tagList"] = getTagList()
	c.Data["profile"] = getProfile()
	c.Data["articleList"] = getArticleListByTagId(tagId)
	c.Data["topTenViewArt"] = getArticleTopThree()

	c.Layout = "home/index.html"
	c.TplName = "home/list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header"] = "home/header.html"
	c.LayoutSections["footer"] = "home/footer.html"
	c.LayoutSections["left"] = "home/left.html"
}

func getArticleList() []models.Article {
	articleList, errArticle := models.GetNormalArticleList()
	if errArticle != nil {
		log.Printf("get article list error: %s", errArticle)
	}
	return articleList
}

func getTagList() []models.ArticleTag {
	tagList, errTag := models.GetTagList()
	if errTag != nil {
		log.Printf("get tag list error: %s", errTag)
	}
	return tagList
}

func getTotalArticleCount() int64 {
	totalArticleCount, errTotal := models.GetTotalArticleCount()
	if errTotal != nil {
		log.Printf("get article count error: %s", errTotal)
	}
	return totalArticleCount
}

func getLinkList() []models.Links {
	linkList, errLinks := models.GetLinkList()
	if errLinks != nil {
		log.Printf("get linkList error: %s", errLinks)
	}
	return linkList
}

func getTopArticle() models.Article {
	topArticle, errTop := models.GetTopArticle()
	if errTop != nil {
		log.Printf("get top article error: %s", errTop)
	}
	return topArticle
}

func getArticleInfoById(id string) models.Article {
	article, err := models.GetArticleInfoById(id)
	if err != nil {
		log.Printf("get article by id error: %s", err)
	}
	return article
}

func getProfile() models.Profile {
	profile, _ := models.GetOneProfile()
	return profile
}

func getArticleListByTagId(tagId string) []models.Article {
	articleList, err := models.GetArticleListByTagId(tagId)
	if err != nil {
		log.Printf("get article by tagid error: %s", err)
	}
	return articleList
}

func updateArticleReadNum(id string) {
	err := models.UpdateArticleReadNum(id)
	if err != nil {
		log.Printf("update article readNum error: %s", err)
	}
}

func getArticleTopThree() []models.Article {
	article, err := models.GetArticleTopThree()
	if err != nil {
		log.Printf("get article top 3 error: %s", err)
	}
	return article
}
