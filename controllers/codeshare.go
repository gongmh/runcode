package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/helpers"
	"github.com/gongmh/runcode/service"
)

type CodeShareController struct {
	beego.Controller
}

func (this *CodeShareController) Get() {
	shareID := this.GetString("id")
	codeMap, lang, codeStat, err := service.GetCodeShareDetail(shareID)
	if err != nil {
		logs.Warn("CodeShareController: GetCodeShareDetail error. err=", err.Error())
		respMap := helpers.GetResponseJson(err)
		this.Data["json"] = respMap
		this.ServeJSON()
	}
	this.Data["default_lag"] = lang
	codeStr, _ := json.Marshal(codeMap)
	this.Data["code_map"] = string(codeStr)
	this.Data["code_stat"] = codeStat
	this.Data["email"] = beego.AppConfig.String("email")
	this.TplName = "tools/index.tpl"
}

func (this *CodeShareController) Post() {
	codeType := this.GetString("code_type")
	codeContent := this.GetString("code_content")

	logs.Info("[Info] _com_request_in codeType=%s, codeContent=%#v", codeType, codeContent)

	if len(codeType) == 0 || len(codeContent) == 0 {
		respMap := helpers.GetResponseJson(helpers.ParamsError)
		this.Data["json"] = respMap
		this.ServeJSON()
		logs.Info("[Info] _com_request_out result=%#v", respMap)
		return
	}

	shareId, err := service.GenCodeShare(codeType, codeContent)
	if err != nil {
		respMap := helpers.GetResponseJson(helpers.ParamsError)
		respMap["data"] = ""
		this.Data["json"] = respMap
		this.ServeJSON()
		logs.Info("[Info] _com_request_out result=%#v", respMap)
		return
	}

	respMap := helpers.GetResponseJson(helpers.Success)
	respMap["data"] = shareId
	this.Data["json"] = respMap

	logs.Info("[Info] _com_request_out result=%#v", respMap)
	this.ServeJSON()
}
