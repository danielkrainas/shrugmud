package storage

import (
	"github.com/danielkrainas/shrugmud/config"
)

var (
	drivers map[string]StorageDriverFactory = make(map[string]StorageDriverFactory)
)

func RegisterDriver(key string, factory StorageDriverFactory) {
	drivers[key] = factory
}

type StorageDriverFactory func() StorageDriver

type StorageDriver interface {
	WriteBlob(key string, data []byte) error
	ReadBlob(key string) ([]byte, error)
	Stat(key string) (*BlobStat, error)
	RemoveBlob(key string) error
}

type BlobStat struct {
	Name string
}

type Storage struct {
	driver StorageDriver
}

func New(storageConfig *config.StorageConfig) *Storage {
	return &Storage{
		driver: drivers[storageConfig.Driver](),
	}
}
