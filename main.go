package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/gongmh/runcode/routers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	beego.Run()
}

func init() {
	// init database
	err := orm.RegisterDataBase(
		beego.AppConfig.String("dbaliasname"),
		beego.AppConfig.String("dbdriver"),
		beego.AppConfig.String("datasource"),
		beego.AppConfig.DefaultInt("dbmaxidleconns", 2),
		beego.AppConfig.DefaultInt("dbmaxopenconns", 10),
	)
	if err != nil {
		panic(err)
	}
}
