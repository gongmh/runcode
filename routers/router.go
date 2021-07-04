package routers

import (
	"github.com/astaxie/beego"
	"github.com/gongmh/runcode/controllers"
)

func init() {
	//代码运行
	beego.Router("/tools/runcode", &controllers.CodeRunController{})
	beego.Router("/tools/s", &controllers.CodeShareController{})
	beego.Router("/tools/v2/runcode", &controllers.CodeRunControllerV2{})

	//地图
	beego.Router("/tools/mapshow", &controllers.MapShow{})
	beego.Router("/tools/dajuesi", &controllers.Dajuesi{})
	beego.Router("/tools/zixishi", &controllers.ZXS{})

	beego.SetStaticPath("/tools/static", "static")
}
