package controllers

import (
	"github.com/danielkrainas/shrugmud/server"
)

type playingState struct {
}

type playingCtrl struct {
	handler server.InputHandler
}

func Playing(handler server.InputHandler) server.Ctrl {
	return &playingCtrl{
		handler: handler,
	}
}

func (ctrl *playingCtrl) Do(input string, d *server.Descriptor) error {
	return nil
}

func (ctrl *playingCtrl) ValidateState(state server.CtrlState) bool {
	_, ok := state.(playingState)
	return ok
}

func (ctrl *playingCtrl) NewState(d *server.Descriptor) server.CtrlState {
	return &playingState{}
}
