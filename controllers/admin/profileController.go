package admin

import (
	"goblog/models"
	"log"
	"net/http"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) List() {

	profileList, num := models.GetProfileList()

	c.Data["num"] = num
	c.Data["statusMap"] = StatusMap
	c.Data["profileList"] = profileList

	c.Layout = "admin/profile/list.html"
	c.TplName = "admin/header.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["nav"] = "admin/nav.html"
}

func (c *ProfileController) Add() {
	if c.IsPost() {
		profile := models.Profile{}
		if err := c.ParseForm(&profile); err != nil {
			c.RespJson(StatZero, "参数解析失败")
		}

		if models.AddProfile(&profile) == 0 {
			c.RespJson(StatZero, "添加失败")
		}
		c.RespJson(StatOne, "添加成功")
	}
	c.Data["statusMap"] = StatusMap
	c.Layout = "admin/profile/add.html"
	c.TplName = "admin/header.html"
}

func (c *ProfileController) Edit() {
	if c.IsPost() {
		id := c.GetString("id")
		nickname := c.GetString("nickname")
		motto := c.GetString("motto")
		status := c.GetString("status")
		if err := models.UpdateProfile(id, nickname, motto, status); err != nil {
			log.Printf("update profile error: %s", err)
			c.RespJson(StatZero, "更新失败")
		}
		c.RespJson(StatOne, "编辑成功")
	} else {
		info, _ := models.GetOneProfile()
		c.Data["info"] = info
		c.Data["statusMap"] = StatusMap
		c.Layout = "admin/profile/edit.html"
		c.TplName = "admin/header.html"
	}
}

func (c *ProfileController)DeleteProfile() {
	id, err := c.GetInt("id")
	if err != nil {
		log.Printf("get id error: %s", err)
	}
	if errProfile := models.DeleteProfileById(id); errProfile != nil {
		log.Printf("delete profile error: %s", errProfile)
	}
	c.Redirect("/profile/list", http.StatusFound)
}
