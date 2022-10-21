package config

import (
	"CaiPu/library/sql"
	"github.com/BurntSushi/toml"
	"log"
)

type Configuration struct {
	Port     string       `toml:"port"`
	Database sql.Database `toml:"database"`
}

var config *Configuration

func Init(configFile string) *Configuration {
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatalf("failed to parse file error(%v)", err)
	}
	return config
}

func Load() *Configuration {
	return config
}
