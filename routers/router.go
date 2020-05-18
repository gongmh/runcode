package routers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/controllers"
)

func init() {
	beego.Router("/tools/runcode", &controllers.CodeRunController{})
	beego.Router("/tools/s", &controllers.CodeShareController{})
	beego.Router("/tools/v2/runcode", &controllers.CodeRunControllerV2{})

	beego.SetStaticPath("/tools/static", "static")
}
