package world

import (
	"github.com/danielkrainas/shrugmud/config"
)

type Realm interface {
	Name() string
}

func New(worldConfig *config.WorldConfig) Realm {
	return &mudRealm{
		name: worldConfig.Name,
	}
}

type mudRealm struct {
	name string
}

func (realm *mudRealm) Name() string {
	return realm.name
}
