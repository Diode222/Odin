package chatMessage

import (
	"fmt"
	"github.com/Diode222/Odin/dao/chatMessage"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"log"
	"net/http"
	"sync"
)

type ChatMessageController struct {
	dao chatMessage.ChatMessageDaoInterface
}

var c *ChatMessageController
var once sync.Once

func NewChatMessageCtrl(dao chatMessage.ChatMessageDaoInterface) *ChatMessageController {
	once.Do(func() {
		c = &ChatMessageController{
			dao: dao,
		}
	})
	return c
}

func (c *ChatMessageController) PutChatMessageList(context *gin.Context) {
	chatMessageList := &pb.ChatMessageList{
		ChatMessages: []*pb.ChatMessage{},
	}
	buf := make([]byte, 4096, 4096)
	n, _ := context.Request.Body.Read(buf)
	if n <= 0 {
		context.Status(http.StatusAccepted)
		context.String(http.StatusAccepted, "status", "No readable chat message data")
		log.Println("No readable chat message data")
		return
	}

	err := proto.Unmarshal(buf[:n], chatMessageList)
	if err != nil {
		log.Println(fmt.Sprintf("chatMessageList unmarshal failed, err: %s", err.Error()))
		context.Error(err)
		return
	}

	status, err := c.dao.PutChatMessageList(context.Request.Context(), chatMessageList)
	if err != nil {
		log.Printf(fmt.Sprintf("PutChatMessageList failed, err: %s", err.Error()))
		context.Error(err)
		context.String(http.StatusNotImplemented, "status", "Put chat message list failed")
	} else {
		if status.GetOK() {
			context.String(http.StatusOK, "status", "OK")
		} else {
			context.String(http.StatusNotImplemented, "status", "Put chat message list failed")
		}
	}
}
