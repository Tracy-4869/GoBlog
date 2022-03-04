package models

import "time"

type ArticleReadNum struct {
	Id         uint      `orm:"column(id); pk; auto; description(主键);"`
	ArticleId  uint      `orm:"column(article_id); unique; description(文章id);"`
	ReadNum    uint      `orm:"column(read_num); default(0); description(阅读量);"`
	Status     uint8     `orm:"column(status); default(1); description(状态 1:正常 2:冻结);"form:"status"`
	CreateTime time.Time `orm:"column(create_time); auto_now_add; type(datetime); description(创建时间);"form:"-"`
}
