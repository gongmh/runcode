package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/service"
)

type CodeRunControllerV2 struct {
	beego.Controller
}

func (this *CodeRunControllerV2) Get() {
	codeID := this.Ctx.GetCookie("tools_code")
	codeMap, lang, _, err := service.GetCodeTemplate(codeID)
	if err != nil {
		logs.Warn("CodeRunControllerV2: GetCodeTemplate error. err=", err.Error())
	}
	this.Data["default_lag"] = lang
	codeStr, _ := json.Marshal(codeMap)
	this.Data["code_map"] = string(codeStr)
	this.TplName = "tools/code_mirror.tpl"
}
