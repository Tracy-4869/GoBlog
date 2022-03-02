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
func (c *Admin) TableEngine() string {
	return "INNODB"
}

// 通过用户名获取用户信息
func GetAdminInfoByName(userName string) (admin Admin, err error) {
	err = orm.NewOrm().QueryTable(&admin).Filter("name", userName).One(&admin)
	return
}

// 通过管理员id获取管理员信息
func GetAdminInfoById(id string) (Admin, error) {
	admin := Admin{}
	err := orm.NewOrm().QueryTable(&admin).Filter("id", id).One(&admin)
	return admin, err
}

// 通过管理员用户名获取管理员等级
func GetAdminGradeByName(userName string) (int, error) {
	admin := Admin{}
	err := orm.NewOrm().QueryTable(&admin).Filter("name", userName).One(&admin)
	return int(admin.Grade), err
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

// 管理员信息更改
func EditAdmin(admin Admin) error {
	_, err := orm.NewOrm().Update(&admin)
	return err
}

// 删除管理员
func DeleteAdminById(id int) error {
	_, err := orm.NewOrm().Delete(&Admin{Id: uint(id)})
	return err
}
