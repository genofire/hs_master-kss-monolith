// Package with the mostly static content (models) of this microservice
package models

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// Config file for this daemon (mor information at the config_example.conf in this git repository)
type Config struct {
	// address under which the api and static content of the webserver runs
	WebserverBind string `toml:"webserver_bind"`

	// path to deliver static content
	Webroot string `toml:"webroot"`

	Database    database.Config   `toml:"database"`
	GoodRelease GoodReleaseConfig `toml:"good_release"`
	CacheClean  CacheWorkerConfig `toml:"cache_clean"`

	// path to the svg image templaes to show availablity of a given good with a traffic light food labeling system
	GoodAvailablityTemplate string `toml:"good_availablity_template"`
	GoodFreshnessTemplate string `toml:"good_freshness_template"`

	// URLs to other microservices that this services uses
	MicroserviceDependencies struct {
		Product    string `toml:"product"`
		Permission string `toml:"permission"`
	} `toml:"microservice_dependencies"`
}

// Configuration of the Worker to clean the cache from values of other microservice
type CacheWorkerConfig struct {
	// Run Worker every Duration
	Every Duration
	// Remove cache, which is not used since Duration
	After Duration
}

// Configuration of the Worker to release locked goods after a time period
type GoodReleaseConfig struct {
	// Run worker every Duration
	Every Duration `toml:"every"`
	// unlock which is not used since Duration
	After Duration `toml:"after"`
}

// Function that reads a config model from a given path of a yml file
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
