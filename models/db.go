package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	mysqlDriver := beego.AppConfig.String("mysqldriver")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlPort := beego.AppConfig.String("mysqlport")
	dataBase := beego.AppConfig.String("database")
	charset := beego.AppConfig.String("charset")

	datasource := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + dataBase + "?charset=" + charset

	orm.RegisterModel(new(Admin), new(ArticleTag), new(Article))

	orm.RegisterDriver(mysqlDriver, orm.DRMySQL)
	orm.RegisterDataBase("default", mysqlDriver, datasource)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
