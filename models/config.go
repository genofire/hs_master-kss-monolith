package models

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

//Config the config File of this daemon
type Config struct {
	WebserverBind string          `toml:"webserver_bind"`
	Database      database.Config `toml:"database"`
	GoodRelease   struct {
		After Duration `toml:"after"`
		Timer Duration `toml:"timer"`
	} `toml:"good_release"`
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
