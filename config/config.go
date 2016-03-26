package config

import (
	"os"
	"strconv"

	"github.com/danielkrainas/shrugmud/logging"
)

type WorldConfig struct {
	Name string
}

type ServerConfig struct {
	Port int
	Host string
}

type StorageConfig struct {
	Driver string
}

type Config struct {
	Storage *StorageConfig
	Server  *ServerConfig
	World   *WorldConfig
}

func newConfig() *Config {
	config := &Config{}
	config.Server = &ServerConfig{}
	config.Server.Port = 7654
	config.Server.Host = "localhost"
	config.World = &WorldConfig{}
	config.World.Name = "ShrugMUD"
	config.Storage = &StorageConfig{}
	config.Storage.Driver = "inmemory"
	return config
}

func validKey(key string, value string) bool {
	if value == "" {
		logging.Trace.Printf("WARNING: key %s should not be empty", key)
		return false
	}

	return true
}

func LoadConfig() (*Config, error) {
	config := newConfig()

	if config.Server != nil {
		strPort := os.Getenv("SHRUG_SERVER_PORT")
		if port, err := strconv.ParseInt(strPort, 10, 32); err == nil {
			config.Server.Port = int(port)
		}

		host := os.Getenv("SHRUG_SERVER_HOST")
		if validKey("SHRUG_SERVER_HOST", host) {
			config.Server.Host = host
		}
	}

	if config.World != nil {
		name := os.Getenv("SHRUG_WORLD_NAME")
		if validKey("SHRUG_WORLD_NAME", name) {
			config.World.Name = name
		}
	}

	if config.Storage != nil {
		driverName := os.Getenv("SHRUG_STORAGE")
		if validKey("SHRUG_STORAGE", driverName) {
			config.Storage.Driver = driverName
		}
	}

	return config, nil
}
