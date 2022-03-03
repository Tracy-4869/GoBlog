package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticleTag struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"form:"-"`
	Name       string    `orm:"column(name); size(32); default(); description(标签名);"form:"name"`
	Total      uint      `orm:"column(total); default(0); description(标签下文章的总数);"form:"total"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}

// 定义表的存储引擎
func (c *ArticleTag) TableEngine() string {
	return "INNODB"
}

// 获取所有标签信息
func GetTagList() ([]ArticleTag, error) {
	tags := []ArticleTag{}
	_, err := orm.NewOrm().QueryTable("article_tag").All(&tags)
	return tags, err
}

// 添加标签
func AddTag(articleTag ArticleTag) error {
	_, err := orm.NewOrm().Insert(&articleTag)
	return err
}

// 根据id获取标签信息
func GetTagInfoById(id string) (ArticleTag, error) {
	articleTag := ArticleTag{}
	err := orm.NewOrm().QueryTable("article_tag").Filter("id", id).One(&articleTag)
	return articleTag, err
}

// 更新标签信息
func UpdateTagById(id, name, status string) error {
	_, err := orm.NewOrm().QueryTable("article_tag").Filter("id", id).Update(orm.Params{
		"name": name,
		"status": status,
	})
	return err
}

// 删除标签
func DeleteTagById(id int) error {
	_, err := orm.NewOrm().Delete(&ArticleTag{Id: uint(id)})
	return err
}

// 该标签下文章数量加1
func AddTotalById(id uint) error {
	_, err := orm.NewOrm().QueryTable("article_tag").Filter("id", id).Update(orm.Params{
		"total": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}
func AddTotalByStringId(id string) error {
	_, err := orm.NewOrm().QueryTable("article_tag").Filter("id", id).Update(orm.Params{
		"total": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

// 该标签下文章数量减1
func MinusTotalById(id string) error {
	_, err := orm.NewOrm().QueryTable("article_tag").Filter("id", id).Update(orm.Params{
		"total": orm.ColValue(orm.ColMinus, 1),
	})
	return err
}

