package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	TagId      uint      `orm:"column(tag_id); default(0); description(标签id);"form:"tag_id"`
	Title      string    `orm:"column(title); size(648); default(); description(文章标题);"form:"title"`
	Abstract   string    `orm:"column(abstract); size(128); default(); description(文章摘要);"form:"abstract"`
	Content    string    `orm:"column(content); type(text); default(); description(文章内容);"form:"content"`
	Author     string    `orm:"column(author); size(32); default(); description(文章作者);"form:"author"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (u *Article) TableEngine() string {
	return "INNODB"
}

// 获取所有文章
func GetArticleList() ([]Article, error) {
	articleList := []Article{}
	_, err := orm.NewOrm().QueryTable("article").All(&articleList)
	return articleList, err
}
