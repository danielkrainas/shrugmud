package controllers

import (
	"github.com/danielkrainas/shrugmud/logging"
	"github.com/danielkrainas/shrugmud/server"
)

type nannyState struct {
}

type nannyCtrl struct {
}

func Nanny() server.Ctrl {
	return &nannyCtrl{}
}

func (ctrl *nannyCtrl) Do(input string, d *server.Descriptor) error {
	logging.Trace.Printf("Ctrl.Nanny: processing %s", input)
	return nil
}

func (ctrl *nannyCtrl) ValidateState(state server.CtrlState) bool {
	_, ok := state.(nannyState)
	return ok
}

func (ctrl *nannyCtrl) NewState(d *server.Descriptor) server.CtrlState {
	return &nannyState{}
}
