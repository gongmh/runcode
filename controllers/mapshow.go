package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type MapShow struct {
	beego.Controller
}

func (this *MapShow) Get() {
	showInfo := service.GenMapShowInfo()
	this.Data["show_info"] = showInfo
	this.TplName = "map/index.tpl"
}
