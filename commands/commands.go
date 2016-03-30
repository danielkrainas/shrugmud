package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/danielkrainas/shrugmud/server"
	//"github.com/danielkrainas/shrugmud/world"
)

type CmdRunner struct {
	inputBuffer chan *CmdContext
	processing  bool
	Packages    map[string]*CmdPackage
}

type CmdContext struct {
	CmdName string
	Args    []string
}

func NewCmdRunner() *CmdRunner {
	return &CmdRunner{
		inputBuffer: make(chan *CmdContext),
		Packages:    make(map[string]*CmdPackage),
		processing:  false,
	}
}

func (runner *CmdRunner) RegisterPackage(pkg *CmdPackage) error {
	_, ok := runner.Packages[pkg.Name]
	if !ok {
		return errors.New(fmt.Sprintf("command package '%s' already registered.", pkg.Name))
	}

	runner.Packages[pkg.Name] = pkg
	return nil
}

func (runner *CmdRunner) InputHandler() server.InputHandler {
	return func(input string, d *server.Descriptor) error {
		parts := strings.Split(input, " ")

		ctx := &CmdContext{
			CmdName: parts[0],
			Args:    parts[1:],
		}

		runner.inputBuffer <- ctx
		runner.Do(ctx)
		return nil
	}
}

func (runner *CmdRunner) Process() {
	runner.processing = true
	go runner.processChannel()
}

func (runner *CmdRunner) StopProcessing() {
	runner.processing = false
}

func (runner *CmdRunner) processChannel() {
	for runner.processing {
		ctx := <-runner.inputBuffer
		runner.Do(ctx)
	}
}

func (runner *CmdRunner) Do(ctx *CmdContext) {
	for _, pkg := range runner.Packages {
		if cmd := pkg.Cmd(ctx.CmdName); cmd != nil {
			cmd.Do(ctx)
			return
		}
	}
}
