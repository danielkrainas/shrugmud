package server

import (
	"fmt"
	//"github.com/danielkrainas/shrugmud/logging"
)

type InputHandler func(input string, d *Descriptor) error

type Ctrl interface {
	Do(input string, d *Descriptor) error
	ValidateState(state CtrlState) bool
	NewState(d *Descriptor) CtrlState
}

type CtrlState interface{}

type CtrlRouter struct {
	routes map[DescState]Ctrl
}

func NewRouter() *CtrlRouter {
	return &CtrlRouter{
		routes: make(map[DescState]Ctrl),
	}
}

func (router *CtrlRouter) Register(ctrl Ctrl, canonicalStates ...string) {
	for _, canonical := range canonicalStates {
		state := canonicalToState(canonical)
		router.routes[state] = ctrl
	}
}

func (router *CtrlRouter) Dispatch(input string, d *Descriptor) error {
	ctrl, found := router.routes[d.State]
	if !found {
		return fmt.Errorf("unsupported state: %s", stateToCanonical(d.State))
	}

	if !ctrl.ValidateState(d.CtrlState) {
		d.CtrlState = ctrl.NewState(d)
	}

	return ctrl.Do(input, d)
}
