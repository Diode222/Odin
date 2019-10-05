package wordFreq

import (
	"fmt"
	"github.com/Diode222/Odin/dao/wordFreq"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type WordFreqController struct {
	dao wordFreq.WordFreqDaoInterface
}

var c *WordFreqController
var once sync.Once

func NewWordFreqCtrl(dao wordFreq.WordFreqDaoInterface) *WordFreqController {
	once.Do(func() {
		c = &WordFreqController{
			dao: dao,
		}
	})
	return c
}

func (c *WordFreqController) GetWordFreqList(context *gin.Context) {
	posStr := context.Query("part_of_speech")
	var pos pb.PartOfSpeech

	switch posStr {
	case "noun":
		pos = pb.PartOfSpeech{Type: pb.PartOfSpeech_NOUN.Enum()}
	case "verb":
		pos = pb.PartOfSpeech{Type: pb.PartOfSpeech_VERB.Enum()}
	case "adjective":
		pos = pb.PartOfSpeech{Type: pb.PartOfSpeech_ADJECTIVE.Enum()}
	case "phrase":
		pos = pb.PartOfSpeech{Type: pb.PartOfSpeech_PHRASE.Enum()}
	default:
		pos = pb.PartOfSpeech{Type: pb.PartOfSpeech_NOUN.Enum()}
	}

	wordFreqList, err := c.dao.GetWordFreqList(context.Request.Context(), pos)
	if err != nil {
		log.Printf(fmt.Sprintf("GetWordFreqList failed, err: %s", err.Error()))
		context.Error(err)
	} else {
		context.ProtoBuf(200, wordFreqList)
	}
}
