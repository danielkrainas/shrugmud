package world

import (
	"fmt"

	"github.com/danielkrainas/shrugmud/config"
	"github.com/danielkrainas/shrugmud/storage"
	"github.com/danielkrainas/shrugmud/world/areas"
	//"github.com/danielkrainas/shrugmud/world/mobiles"
	"github.com/danielkrainas/shrugmud/world/players"
)

type Realm interface {
	Name() string
}

func New(worldConfig *config.WorldConfig, store *storage.Storage) Realm {
	return &mudRealm{
		name:    worldConfig.Name,
		areas:   []*areas.Area{areas.NewLimbo()},
		players: make(map[string]*players.Player),
		store:   store,
	}
}

type mudRealm struct {
	name    string
	store   *storage.Storage
	areas   []*areas.Area
	players map[string]*players.Player
}

func (realm *mudRealm) Name() string {
	return realm.name
}

func (realm *mudRealm) LoadPC(name string) (*players.Player, error) {
	_, found := realm.players[name]
	if found {
		return nil, fmt.Errorf("player already loaded: %s", name)
	}

	return players.Load(name, realm.store)
}
