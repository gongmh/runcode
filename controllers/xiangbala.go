package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type XiangBaLa struct {
	beego.Controller
}

func (this *XiangBaLa) Get() {
	showInfo := service.GenXBLShowInfo()
	this.Data["show_info"] = showInfo
	this.TplName = "map/dajuesi.tpl"
}
