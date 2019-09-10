package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Odin/controllers:GetController"] = append(beego.GlobalControllerRouter["Odin/controllers:GetController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/get`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odin/controllers:NounGetController"] = append(beego.GlobalControllerRouter["Odin/controllers:NounGetController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/get`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
