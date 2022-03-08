package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Links struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(友情名);"form:"name"`
	Url        string    `orm:"column(url); size(64); default(); description(链接地址);"form:"url"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (c *Links) TableEngine() string {
	return "INNODB"
}

//通过id来获取链接信息
func GetLinkInfoById(id string) (Links, error) {
	link := Links{}
	err := orm.NewOrm().QueryTable("links").Filter("id", id).One(&link)
	return link, err
}

// 更新链接信息
func UpdateLinkById(id, name, url, status string) error {
	_, err := orm.NewOrm().QueryTable(new(Links)).Filter("id", id).Update(orm.Params{
		"name":   name,
		"url":    url,
		"status": status,
	})
	return err
}

//获取所有链接
func GetLinkList() ([]Links, error) {
	linksList := []Links{}
	_, err := orm.NewOrm().QueryTable("links").All(&linksList)
	return linksList, err
}

//添加链接
func AddLinkInfo(link Links) error {
	_, err := orm.NewOrm().Insert(&link)
	return err
}

// 删除链接
func DeleteLinkById(id int) error {
	_, err := orm.NewOrm().Delete(&Links{Id: uint(id)})
	return err
}


