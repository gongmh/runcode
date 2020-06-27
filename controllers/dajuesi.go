package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type Dajuesi struct {
	beego.Controller
}

func (this *Dajuesi) Get() {
	//showInfo := service.GenMapShowInfo()
	showInfo := service.GenDajuesiShowInfo()
	this.Data["show_info"] = showInfo
	this.TplName = "map/dajuesi.tpl"
}
