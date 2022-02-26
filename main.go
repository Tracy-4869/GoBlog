package main

import (
	"goblog/models"
	_ "goblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlDriver := beego.AppConfig.String("mysqldriver")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlPort := beego.AppConfig.String("mysqlport")
	dataBase := beego.AppConfig.String("database")
	charset := beego.AppConfig.String("charset")

	datasource := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + dataBase + "?charset=" + charset

	orm.RegisterModel(new(models.Admin))

	orm.RegisterDriver(mysqlDriver, orm.DRMySQL)
	orm.RegisterDataBase("default", mysqlDriver, datasource)
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	beego.Run()
}
