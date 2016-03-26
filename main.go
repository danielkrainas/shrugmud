package main

import (
	"github.com/danielkrainas/shrugmud/config"
	//"github.com/danielkrainas/shrugmud/logging"
	"github.com/danielkrainas/shrugmud/server"
	"github.com/danielkrainas/shrugmud/world"
)

func main() {
	gameConfig := config.LoadConfig()
	realm := world.New(gameConfig.World)
	mudServer := server.New(gameConfig.Server)
	mudServer.Start(realm)
}
