package wordFreq

import (
	"fmt"
	"github.com/Diode222/Odin/dao/wordFreq"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
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

func (ctrl *WordFreqController) GetWordFreqList(context *gin.Context) {
	posStr := context.Query("part_of_speech")
	var pos pb.PartOfSpeech_POSType

	switch posStr {
	case "noun":
		pos = pb.PartOfSpeech_NOUN
	case "verb":
		pos = pb.PartOfSpeech_VERB
	case "adjective":
		pos = pb.PartOfSpeech_ADJECTIVE
	case "phrase":
		pos = pb.PartOfSpeech_PHRASE
	default:
		pos = pb.PartOfSpeech_NOUN
	}

	if (pos == pb.PartOfSpeech_NOUN) {
		buf := make([]byte, 4096, 4096)
		n, err := context.Request.Body.Read(buf)
		fmt.Println("request method: ", context.Request.Method)

		fmt.Println("read data length: ", n)
		if context.Error(err) != nil {
			log.Printf("Read request failed")
		}

		if n > 0 {
			var chatMessageList *pb.ChatMessageList = &pb.ChatMessageList{
				ChatMessages: []*pb.ChatMessage{},
			}
			err = proto.Unmarshal(buf[:n], chatMessageList)
			if err != nil {
				fmt.Println("proto unmarshal failed")
			}
			for _, chatMessage := range chatMessageList.ChatMessages {
				fmt.Println("chatMessage: " + chatMessage.GetMessage())
				fmt.Println("time: " + string(chatMessage.GetTime()))
				fmt.Println("chatPerson: " + chatMessage.GetChatPerson())
			}
		}

		wordFreqList := &pb.WordFreqList{
			WordFreqs: []*pb.WordFreq{},
		}
		wordFreq := &pb.WordFreq{
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

	//data, err := ctrl.dao.GetData(context, pos)
	//
	//if context.Error(err) != nil {
	//	log.Printf(err.Error())
	//} else {
	//	context.JSON(200, view.FormatWordItems(&data))
	//}
}
