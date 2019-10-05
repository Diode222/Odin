package ctrl

import (
	"github.com/Diode222/Odin/app/dao"
	"github.com/Diode222/Odin/controller/chatMessage"
	"github.com/Diode222/Odin/controller/wordFreq"
)

var WordFreqCtrl *wordFreq.WordFreqController = wordFreq.NewWordFreqCtrl(dao.WordFreqDao)

var ChatMessageCtrl *chatMessage.ChatMessageController = chatMessage.NewChatMessageCtrl(dao.ChatMessageDao)
