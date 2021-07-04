package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type ZXS struct {
	beego.Controller
}

func (this *ZXS) Get() {
	showInfo := service.GetZXSCoord()
	this.Data["show_info"] = showInfo
	this.TplName = "map/zxs.tpl"
}
