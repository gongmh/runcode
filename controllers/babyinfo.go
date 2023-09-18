package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type BabyInfo struct {
	beego.Controller
}

func (this *BabyInfo) Get() {
	this.Data["weight_info"] = service.GetChildWeightInfo()
	this.Data["event_info"] = service.GetChildEventInfo()
	this.TplName = "echarts/babyinfo.tpl"
}
