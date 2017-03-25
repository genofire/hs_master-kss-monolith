package models

import (
	"io/ioutil"

	"github.com/influxdata/toml"
)

//Config the config File of this daemon
type Config struct {
	WebserverBind string
}

// ReadConfigFile reads a config model from path of a yml file
func ReadConfigFile(path string) *Config {
	config := &Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(file, config); err != nil {
		panic(err)
	}

	return config
}
