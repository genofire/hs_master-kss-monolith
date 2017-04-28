package models

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

//config file of this daemon (for more the config_example.conf in git repository)
type Config struct {
	// address on which the api and static content webserver runs
	WebserverBind string `toml:"webserver_bind"`

	// path to deliver static content
	Webroot string `toml:"webroot"`

	Database    database.Config   `toml:"database"`
	GoodRelease GoodReleaseConfig `toml:"good_release"`
	CacheClean  CacheWorkerConfig `toml:"cache_clean"`

	// path to the svg image templaes to show availablity of a given good
	GoodAvailablityTemplate string `toml:"good_availablity_template"`

	// URLs to other microservices
	MicroserviceDependencies struct {
		Product    string `toml:"product"`
		Permission string `toml:"permission"`
	} `toml:"microservice_dependencies"`
}

//config of worker to clean caches from values of other microservice
type CacheWorkerConfig struct {
	// run worker every ...
	Every Duration
	// remove cache which is not used since ..
	After Duration
}

//config of worker to release looked goods after a time
type GoodReleaseConfig struct {
	// run worker every ...
	Every Duration `toml:"every"`
	// unlock which is not used since ..
	After Duration `toml:"after"`
}

//reads a config model from path of a yml file
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
