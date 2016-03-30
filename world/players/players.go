package players

import (
	"github.com/danielkrainas/shrugmud/storage"
	"github.com/danielkrainas/shrugmud/world/mobiles"
)

type Player struct {
	Name string
	Ch   *mobiles.Mobile
}

func Load(name string, store *storage.Storage) (*Player, error) {
	return &Player{
		Name: name,
	}, nil
}
