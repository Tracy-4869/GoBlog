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
	ReadNum    uint      `orm:"column(read_num); default(0); description(阅读量);"`
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

// 添加文章
func AddArticle(article Article) error {
	_, err := orm.NewOrm().Insert(&article)
	return err
}

// 通过id获取文章信息
func GetArticleInfoById(id string) (Article, error) {
	article := Article{}
	err := orm.NewOrm().QueryTable("article").Filter("id", id).One(&article)
	return article, err
}

// 更新文章信息
func UpdateArticleById(id, tag_id, title, content, author, status string) error {
	_, err := orm.NewOrm().QueryTable(new(Article)).Filter("id", id).Update(orm.Params{
		"tag_id":  tag_id,
		"title":   title,
		"content": content,
		"author":  author,
		"status":  status,
	})
	return err
}

// 删除文章
func DeleteArticleById(id uint) error {
	_, err := orm.NewOrm().Delete(&Article{Id: id})
	return err
}

// 获取正常状态下的文章总数
func GetTotalArticleCount() (int64, error) {
	count, err := orm.NewOrm().QueryTable(new(Article)).Filter("status", 1).Count()
	return count, err
}
