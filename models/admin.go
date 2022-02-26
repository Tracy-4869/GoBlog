package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(管理员姓名);"form:"name"`
	PassWord   string    `orm:"column(password); size(32); deafult(); description(密码);"form:"password"`
	Grade      uint8     `orm:"column(grade); default(2); description(管理员等级 1:超管 2:普通);"form:"grade"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *Admin) TableEngine() string {
	return "INNODB"
}

// 通过用户名获取用户信息
func GetAdminInfoByName(userName string) (admin Admin, err error) {
	err = orm.NewOrm().QueryTable(&admin).Filter("name", userName).Filter("status", 1).One(&admin)
	return
}

// 获取所有管理员信息
func GetAdminList()([]Admin, error) {
	adminList := []Admin{}
	_, err := orm.NewOrm().QueryTable("admin").All(&adminList)
	return adminList, err
}

// 添加管理员
func AddAdminInfo(admin Admin) error {
	_, err := orm.NewOrm().Insert(&admin)
	return err
}
