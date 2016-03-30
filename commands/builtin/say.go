package builtin

import (
	"github.com/danielkrainas/shrugmud/commands"
)

var (
	doSay = commands.NewCmd("say", doSayHandler)
)

func doSayHandler(ctx *commands.CmdContext) error {
	return nil
}
