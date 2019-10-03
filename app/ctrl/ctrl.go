package ctrl

import (
	"github.com/Diode222/Odin/app/dao"
	"github.com/Diode222/Odin/controller/get"
)

var GetCtrl *get.GetController = get.NewGetCtrl(dao.GetDao)
