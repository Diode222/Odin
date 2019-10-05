package app

import (
	"github.com/Diode222/Odin/app/ctrl"
	"github.com/Diode222/Odin/controller/chatMessage"
	"github.com/Diode222/Odin/controller/wordFreq"
	"sync"
)

type handler struct {
	WordFreqCtrl *wordFreq.WordFreqController
	ChatMessageCtrl *chatMessage.ChatMessageController
}

var h *handler
var once sync.Once

func NewHandler() *handler {
	once.Do(func() {
		h = &handler{
			WordFreqCtrl: ctrl.WordFreqCtrl,
			ChatMessageCtrl: ctrl.ChatMessageCtrl,
		}
	})
	return h
}
