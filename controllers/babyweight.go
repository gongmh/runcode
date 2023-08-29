package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type BabyWeight struct {
	beego.Controller
}

func (this *BabyWeight) Get() {
	showInfo := service.GetChildWeightInfo()
	this.Data["show_info"] = showInfo
	this.TplName = "echarts/babyweight.tpl"
}
