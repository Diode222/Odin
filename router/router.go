package router

import (
	"Odin/controller/get"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&get.NounGetController{})
}
