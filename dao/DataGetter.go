package dao

import (
	"Odin/model"
	"github.com/astaxie/beego/context"
)

type DataGetter interface {
	GetData(ctx *context.Context, dataAmountType model.DataAmountType, pos model.PartOfSpeech) (model.WordItems, error)
}
