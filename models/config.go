package models

import (
	"io/ioutil"

	"github.com/influxdata/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/lib/database"
)

//Config the config File of this daemon
type Config struct {
	WebserverBind string
	Database database.Config
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
