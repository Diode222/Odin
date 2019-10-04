package main

import (
	"fmt"
	"github.com/Diode222/Odin/app"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	handler := app.NewHandler()
	engine.Any("/get", handler.GetCtrl.Get)
	err := engine.Run("172.27.0.15:38888")
	if err != nil {
		log.Printf(fmt.Sprintf("Odin server can not run to listen port: %d, err: %s", 38888, err.Error()))
	}
}
