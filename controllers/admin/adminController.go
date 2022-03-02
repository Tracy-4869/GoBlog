package admin

import (
	"goblog/models"
	"goblog/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AdminController struct {
	BaseController
}

type AdminList struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Grade string `json"grade"`
	Status string `json:"status"`
	CreateTime time.Time `json:"createTime"`
}

// 管理员列表
func (c *AdminController) List() {
	adminList, err := models.GetAdminList()
	if err != nil {
		log.Printf("get adminList error:%s", err)
	}
	c.Data["statusMap"] = StatusMap
	c.Data["adminLevelMap"] = adminLevelMap
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
		if err := c.ParseForm(&admin); err != nil {
			log.Printf("parseform admin error: %s", err)
			c.RespJson(StatZero, "parseform admin error")
		}
		// 管理员权限判断
		if admin.Grade == 1 {
			userName := c.GetSession("userName")
			adminGrade, err := models.GetAdminGradeByName(userName.(string))
			if err != nil {
				log.Printf("get admin grade error: %s", err)
			}
			if adminGrade == 2 {
				c.RespJson(StatZero, "you cannot add superAdmin")
			}
		}
		admin.PassWord = utils.Md5Text(admin.PassWord)
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

// 管理员编辑
func (c *AdminController) Edit() {
	if c.IsPost() {
		id := c.GetString("id")
		status := c.GetString("status")
		grade := c.GetString("grade")
		// 参数验证
		if id == "" || grade == "" || status == "" {
			c.RespJson(StatZero, "参数错误")
		}
		admin, err := models.GetAdminInfoById(id)
		if err != nil {
			log.Printf("get admininfo error: %s", err)
		}
		statusAdmin, errStatus := strconv.Atoi(status)
		if errStatus != nil {
			log.Printf("string to int error: %s", errStatus)
		}
		gradeAdmin, errGrade := strconv.Atoi(grade)
		if errGrade != nil {
			log.Printf("string to int error: %s", errGrade)
		}
		// 管理员权限判断
		if gradeAdmin == 1 {
			userName := c.GetSession("userName")
			adminGrade, err := models.GetAdminGradeByName(userName.(string))
			if err != nil {
				log.Printf("get admin grade error: %s", err)
			}
			if adminGrade == 2 {
				c.RespJson(StatZero, "you cannot add superAdmin")
			}
		}
		admin.Status = uint8(statusAdmin)
		admin.Grade = uint8(gradeAdmin)
		if errUpdate := models.EditAdmin(admin); errUpdate != nil {
			log.Printf("update admin info error: %s", errUpdate)
		}
		c.RespJson(StatOne, "编辑成功")
	} else {
		getId := c.GetString("id")
		info, _ := models.GetAdminInfoById(getId)

		c.Data["info"] = info
		c.Data["adminLevelMap"] = adminLevelMap
		c.Data["statusMap"] = StatusMap

		c.Layout = "admin/admin/edit.html"
		c.TplName = "admin/header.html"
	}
}

// 管理员删除
func (c *AdminController) DeleteAdmin() {
	id, err := c.GetInt("id")
	if err != nil {
		log.Printf("get id error: %s", err)
	}
	// 管理员权限判断
	userName := c.GetSession("userName")
	adminGrade, errGrade := models.GetAdminGradeByName(userName.(string))
	if errGrade != nil {
		log.Printf("get admin grade error:%s", errGrade)
	}
	if adminGrade == 1 {
		if deleteErr := models.DeleteAdminById(id); deleteErr != nil {
			log.Printf("delete admin by id:%d  error: %s", id, deleteErr)
		}
	}
	c.Redirect("/admin/list", http.StatusFound)
}
