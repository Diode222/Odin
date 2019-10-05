package chatMessage

import (
	"github.com/Diode222/Odin/dao/chatMessage"
	"github.com/gin-gonic/gin"
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
	// TODO
}
