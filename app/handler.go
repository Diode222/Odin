package app

import (
	"github.com/Diode222/Odin/app/ctrl"
	"github.com/Diode222/Odin/controller/get"
	"sync"
)

type handler struct {
	GetCtrl *get.GetController
}

var h *handler
var once sync.Once

// TODO
func NewHandler() *handler {
	once.Do(func() {
		h = &handler{
			GetCtrl: ctrl.GetCtrl,
		}
	})
	return h
}
