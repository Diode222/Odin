package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Odin/controller:NounGetController"] = append(beego.GlobalControllerRouter["Odin/controller:NounGetController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/get`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
