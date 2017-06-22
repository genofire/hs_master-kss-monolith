// Package with the mostly static content (models) of this microservice
package models

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// Config file for this daemon (more information at the config_example.conf in this git repository)
type Config struct {
	// address under which the api and static content of the webserver runs
	WebserverBind string `toml:"webserver_bind"`

	// path to deliver static content
	Webroot string `toml:"webroot"`

	Database    database.Config   `toml:"database"`
	GoodRelease GoodReleaseConfig `toml:"good_release"`
	CacheClean  CacheWorkerConfig `toml:"cache_clean"`

	// path to the SVG image templates to show the availability and freshness
	// of a given good with a traffic light food labeling system
	GoodAvailabilityTemplate string `toml:"good_availablity_template"`
	GoodFreshnessTemplate    string `toml:"good_freshness_template"`

	FouledDeleter Duration `toml:"fouled_deleted"`

	// URLs to other microservices, which this service uses
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
	// Unlock those, which are not used since Duration
	After Duration `toml:"after"`
}

// Function that reads a config model from a given path of a .yml file
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
