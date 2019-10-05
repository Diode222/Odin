package main

import (
	"fmt"
	"github.com/Diode222/Odin/app"
	"github.com/Diode222/Odin/app/conf"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	handler := app.NewHandler()
	engine.Any("/wordfreq", handler.WordFreqCtrl.GetWordFreqList)
	engine.Any("/chatmessage", handler.ChatMessageCtrl.PutChatMessageList)
	err := engine.Run(fmt.Sprintf("%s:%d", conf.SERVICE_IP, conf.SERVICE_PORT))
	//err := engine.Run("127.0.0.1:38888")
	if err != nil {
		log.Printf(fmt.Sprintf("Odin server can not run to listen port: %d, err: %s", 38888, err.Error()))
	}
}
