package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	NickName   string    `orm:"column(nickname); size(32); default(); description(昵称);"form:"nickname"`
	Motto      string    `orm:"column(motto); size(64); default(); description(座右铭);"form:"motto"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (p *Profile) TableEngine() string {
	return "INNODB"
}

// 获取个人简介
func GetProfileList() (profile []Profile, num int64) {
	profile = []Profile{}
	num, _ = orm.NewOrm().QueryTable(new(Profile)).All(&profile)
	return
}

func GetOneProfile() (profile Profile, num int64) {
	profile = Profile{}
	num, _ = orm.NewOrm().QueryTable(new(Profile)).All(&profile)
	return
}

// 添加简介
func AddProfile(p *Profile) int64 {
	addRow, err := orm.NewOrm().Insert(p)
	if err != nil {
		log.Printf("insert profile error: %s", err)
		addRow = 0
	}
	return addRow
}

// 更改个人简介
func UpdateProfile(id, nickname, motto, status string) error {
	_, err := orm.NewOrm().QueryTable(new(Profile)).Filter("id", id).Update(orm.Params{
		"nickname": nickname,
		"motto":    motto,
		"status":   status,
	})
	return err
}

// 删除个人简介
func DeleteProfileById(id int) error {
	_, err := orm.NewOrm().Delete(&Profile{Id: uint(id)})
	return err
}
