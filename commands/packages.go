package commands

import ()

type CmdFunc func(ctx *CmdContext) error

type CmdPackage struct {
	Name string
	Cmds map[string]*Cmd
}

type Cmd struct {
	Name string
	do   CmdFunc
}

func (cmd *Cmd) Do(ctx *CmdContext) {
	_ = cmd.do(ctx)
	return
}

func NewCmd(name string, handler CmdFunc) *Cmd {
	return &Cmd{
		Name: name,
		do:   handler,
	}
}

func NewCmdPackage(name string) *CmdPackage {
	return &CmdPackage{
		Name: name,
		Cmds: make(map[string]*Cmd),
	}
}

func (pkg *CmdPackage) Register(cmd *Cmd) {
	pkg.Cmds[cmd.Name] = cmd
}

func (pkg *CmdPackage) Cmd(name string) *Cmd {
	cmd, found := pkg.Cmds[name]
	if !found {
		return nil
	}

	return cmd
}
