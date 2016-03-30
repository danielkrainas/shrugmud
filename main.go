package main

import (
	//"github.com/danielkrainas/shrugmud/logging"

	"github.com/danielkrainas/shrugmud/commands"
	"github.com/danielkrainas/shrugmud/commands/builtin"
	"github.com/danielkrainas/shrugmud/config"
	"github.com/danielkrainas/shrugmud/logging"
	"github.com/danielkrainas/shrugmud/server"
	"github.com/danielkrainas/shrugmud/server/controllers"
	"github.com/danielkrainas/shrugmud/storage"
	"github.com/danielkrainas/shrugmud/world"
)

func main() {
	gameConfig, err := config.LoadConfig()
	if err != nil {
		logging.Error.Fatal(err)
	}

	store := storage.New(gameConfig.Storage)

	cmds := commands.NewCmdRunner()
	cmds.RegisterPackage(builtin.Package)

	router := server.NewRouter()
	router.Register(controllers.Nanny(), "welcome", "login")
	router.Register(controllers.Playing(cmds.InputHandler()), "playing")

	realm := world.New(gameConfig.World, store)
	mudServer := server.New(gameConfig.Server, router)

	cmds.Process()
	mudServer.Start(realm)
}
