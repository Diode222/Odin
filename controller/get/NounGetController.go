package get

import (
	"Odin/dao"
	"Odin/model"
	"Odin/view"
	"github.com/astaxie/beego"
)

type NounGetController struct {
	beego.Controller
	dao.DataGetter
}

// @router /get [Get]
func (ctrl *NounGetController) Get() {
	dataAmount := ctrl.GetString("data_amount")
	if dataAmount == "" {
		dataAmount = "intense"
	}

	var data model.WordItems
	var err error
	switch dataAmount {
	case "intense":
		data, err = ctrl.DataGetter.GetData(ctrl.Ctx, model.INTENSIVE_DATA, model.NOUN)
	case "moderate":
		data, err = ctrl.DataGetter.GetData(ctrl.Ctx, model.MODERATE_DATA, model.NOUN)
	case "sparse":
		data, err = ctrl.DataGetter.GetData(ctrl.Ctx, model.SPARSE_DATA, model.NOUN)
	default:
		data, err = ctrl.DataGetter.GetData(ctrl.Ctx, model.INTENSIVE_DATA, model.NOUN)
	}
	if (err != nil) {
		ctrl.Data["json"] = err.Error()
	} else {
		ctrl.Data["json"] = view.FormatWordItems(data)
	}
}
