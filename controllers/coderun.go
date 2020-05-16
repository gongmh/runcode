package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gongmh/runcode/helpers"
	"github.com/gongmh/runcode/service"
)

type CodeRunController struct {
	beego.Controller
}

func (this *CodeRunController) Get() {
	codeID := this.Ctx.GetCookie("tools_code")
	codeMap, lang, codeStat, err := service.GetCodeTemplate(codeID)
	if err != nil {
		logs.Warn("CodeRunController: GetCodeTemplate error. err=", err.Error())
	}

	this.Data["default_lag"] = lang
	codeStr, _ := json.Marshal(codeMap)
	this.Data["code_map"] = string(codeStr)
	this.Data["code_stat"] = codeStat
	this.Data["email"] = beego.AppConfig.String("email")
	this.TplName = "tools/index.tpl"
}

func (this *CodeRunController) Post() {
	codeType := this.GetString("code_type")
	codeContent := this.GetString("code_content")
	clientIp := this.Ctx.Input.IP()
	userAgent := this.Ctx.Input.UserAgent()
	beego.BeeLogger.Info("[Info] _com_request_in codetype=%s, code_content=%#v", codeType, codeContent)

	if len(codeType) == 0 || len(codeContent) == 0 {
		respMap := helpers.GetResponseJson(helpers.ParamsError)
		this.Data["json"] = respMap
		this.ServeJSON()
		beego.BeeLogger.Info("[Info] _com_request_out result=%#v", respMap)
		return
	}

	s, codeId, err := service.CodeRun(codeType, codeContent, clientIp, userAgent)
	if err != nil {
		respMap := helpers.GetResponseJson(helpers.ParamsError)
		respMap["data"] = s
		this.Data["json"] = respMap
		this.ServeJSON()
		beego.BeeLogger.Info("[Info] _com_request_out result=%#v", respMap)
		return
	}

	respMap := helpers.GetResponseJson(helpers.Success)
	respMap["data"] = s
	this.Data["json"] = respMap

	this.Ctx.SetCookie("tools_code", codeId, 86400*7, "/")
	beego.BeeLogger.Info("[Info] _com_request_out result=%#v", respMap)
	this.ServeJSON()
}
