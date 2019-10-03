package dao

import (
	"github.com/Diode222/Odin/model"
	"github.com/gin-gonic/gin"
)

type DataGetter interface {
	GetData(ctx *gin.Context, dataAmountType model.DataAmountType, pos model.PartOfSpeech) (model.WordItems, error)
}
