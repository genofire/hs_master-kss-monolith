package models

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

//Config the config File of this daemon
type Config struct {
	WebserverBind            string            `toml:"webserver_bind"`
	Webroot                  string            `toml:"webroot"`
	Database                 database.Config   `toml:"database"`
	GoodRelease              GoodReleaseConfig `toml:"good_release"`
	CacheClean               CacheWorkerConfig `toml:"cache_clean"`
	GoodAvailablityTemplate  string            `toml:"good_availablity_template"`
	MicroserviceDependencies struct {
		Product    string `toml:"product"`
		Permission string `toml:"permission"`
	} `toml:"microservice_dependencies"`
}

type CacheWorkerConfig struct {
	Every Duration
	After Duration
}

type GoodReleaseConfig struct {
	After Duration `toml:"after"`
	Every Duration `toml:"every"`
}

// ReadConfigFile reads a config model from path of a yml file
func ReadConfigFile(path string) *Config {
	config := &Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Log.Panic(err)
	}

	if err := toml.Unmarshal(file, config); err != nil {
		log.Log.Panic(err)
	}

	return config
}
