package drivers

import (
	"github.com/danielkrainas/shrugmud/storage"
)

func init() {
	storage.RegisterDriver("memory", memoryStorageFactory)
}

func memoryStorageFactory() storage.StorageDriver {
	return memoryStorage{}
}

type memoryStorage struct {
}

func (driver *fsStorageDriver) ReadBlob(key string) ([]byte, error) {

}

func (driver *fsStorageDriver) WriteBlob(key string, data []byte) error {

}

func (driver *fsStorageDriver) Stat(key string) (*BlobStat, error) {

}

func (driver *fsStorageDriver) RemoveBlob(key string) error {

}
