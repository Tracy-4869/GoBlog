package main

import (
	"goblog/models"
	"goblog/routers"

	"github.com/astaxie/beego"
	
)

func main() {
	models.Init()
	routers.Init()
	beego.Run()
}
