package builtin

import (
	"github.com/danielkrainas/shrugmud/commands"
)

var (
	Package *commands.CmdPackage
)

func init() {
	Package = commands.NewCmdPackage("builtin")
	Package.Register(doSay)
}
