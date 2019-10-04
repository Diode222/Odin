package get

import (
	"fmt"
	"github.com/Diode222/Odin/dao"
	"github.com/Diode222/Odin/model"
	"github.com/Diode222/Odin/proto_gen"
	"github.com/Diode222/Odin/view"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
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
	if (pos == model.NOUN) {
		buf := make([]byte, 4096, 4096)
		n, err := context.Request.Body.Read(buf)
		if err != nil {
			log.Printf("Read request failed")
		}

		var chatMessageList *proto_gen.ChatMessageList
		err = proto.Unmarshal(buf[:n], chatMessageList)
		if err != nil {
			fmt.Println("proto unmarshal failed")
		}
		for _, chatMessage := range chatMessageList.ChatMessages {
			fmt.Println("chatMessage: " + chatMessage.GetMessage())
			fmt.Println("time: " + string(chatMessage.GetTime()))
			fmt.Println("chatPerson: " + chatMessage.GetChatPerson())
		}

		wordFreqList := &proto_gen.WordFreqList{
			WordFreqs: []*proto_gen.WordFreq{},
		}
		var wordFreq *proto_gen.WordFreq
		wordFreq = &proto_gen.WordFreq{
			Word: proto.String("车不车"),
			Count: proto.Int(50),
		}

		wordFreqList.WordFreqs = append(wordFreqList.WordFreqs, wordFreq)

		responseData, err := proto.Marshal(wordFreqList)
		if err != nil {
			fmt.Println("proto.Marshal failed")
		}
		context.Data(200, "OK", responseData)

		return
	}
	if context.Error(err) != nil {
		log.Printf(err.Error())
	} else {
		context.JSON(200, view.FormatWordItems(&data))
	}
}
