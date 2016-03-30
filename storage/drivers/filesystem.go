package drivers

import (
	"github.com/danielkrainas/shrugmud/storage"
)

func init() {
	storage.RegisterDriver("filesystem", fsStorageFactory)
}

func fsStorageFactory() storage.StorageDriver {
	return fsStorageDriver{}
}

type fsStorageDriver struct {
}

func (driver *fsStorageDriver) ReadBlob(key string) ([]byte, error) {

}

func (driver *fsStorageDriver) WriteBlob(key string, data []byte) error {

}

func (driver *fsStorageDriver) Stat(key string) (*BlobStat, error) {

}

func (driver *fsStorageDriver) RemoveBlob(key string) error {

}
