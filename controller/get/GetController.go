package get

import (
	"github.com/Diode222/Odin/dao"
	"github.com/Diode222/Odin/model"
	"github.com/Diode222/Odin/view"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type GetController struct {
	dao dao.DataGetter
}

var c *GetController
var once sync.Once

func NewGetCtrl(getDAO dao.DataGetter) *GetController {
	once.Do(func() {
		c = &GetController{
			dao: getDAO,
		}
	})
	return c
}

// @router /get [Get]
func (ctrl *GetController) Get(context *gin.Context) {
	dataAmountStr := context.Query("data_amount")
	posStr := context.Query("part_of_speech")
	var dataAmount model.DataAmountType
	var pos model.PartOfSpeech

	switch dataAmountStr {
	case "intense":
		dataAmount = model.INTENSIVE_DATA
	case "moderate":
		dataAmount = model.MODERATE_DATA
	case "sparse":
		dataAmount = model.SPARSE_DATA
	default:
		dataAmount = model.INTENSIVE_DATA
	}

	switch posStr {
	case "noun":
		pos = model.NOUN
	case "verb":
		pos = model.VERB
	case "adjective":
		pos = model.ADJECTIVE
	case "phrase":
		pos = model.PHRASE
	default:
		pos = model.NOUN
	}

	data, err := ctrl.dao.GetData(context, dataAmount, pos)
	if context.Error(err) != nil {
		log.Printf(err.Error())
	} else {
		context.JSON(200, view.FormatWordItems(&data))
	}
}
