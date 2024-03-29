package dao

import (
	"github.com/Diode222/Odin/dao/chatMessage"
	"github.com/Diode222/Odin/dao/wordFreq"
	"github.com/Diode222/Odin/dao/wordSplit"
)

var WordFreqDao wordFreq.WordFreqDaoInterface = wordFreq.NewWordFreqDao()

var ChatMessageDao chatMessage.ChatMessageDaoInterface = chatMessage.NewChatMessageDao()

var WordSplitDao wordSplit.WordSplitDaoInterface = wordSplit.NewWordSplitDao()
