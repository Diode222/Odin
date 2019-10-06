package wordFreq

import (
	"fmt"
	"github.com/Diode222/Odin/dao/wordFreq"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	//list := &pb.ChatMessageList{
	//	ChatMessages:         []*pb.ChatMessage{},
	//}
	//msg1 := "名词是一种很好的词语"
	//var time1 int64 = 2345234
	//person1 := "文件聊天助手"
	//list.ChatMessages = append(list.ChatMessages, &pb.ChatMessage{
	//	Message:              &msg1,
	//	Time:                 &time1,
	//	ChatPerson:           &person1,
	//	WordAndPosList:       nil,
	//})
	//l, e := client.GetWordSplitClient().GetWordSplittedMessageList(context.Request.Context(), list)
	//if e != nil {
	//	fmt.Println("test mimiron get failed, err: ", e.Error())
	//}
	//fmt.Println("xixi:" + string(len(l.GetChatMessages())))
	//if len(l.GetChatMessages()) > 0 {
	//	for _, chatMessage := range l.GetChatMessages() {
	//		for _, x := range chatMessage.GetWordAndPosList() {
	//			fmt.Println(x.GetWord())
	//			fmt.Println(x.GetPos().GetType())
	//		}
	//	}
	//}

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
		context.String(http.StatusNotImplemented, "status", "Get word frequence list failed")
	} else {
		log.Println("wordFreqList length: ", len(wordFreqList.WordFreqs))
		context.ProtoBuf(http.StatusOK, wordFreqList)
	}
}
