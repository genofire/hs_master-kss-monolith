package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	Write  *gorm.DB
	Read   *gorm.DB
	config *Config
	models []interface{}
)

type Config struct {
	Type           string
	Connection     string
	ReadConnection string
}

func Open(c Config) (err error) {
	config = &c
	Write, err = gorm.Open(config.Type, config.Connection)
	if err != nil {
		return
	}
	if len(config.ReadConnection) > 0 {
		Read, err = gorm.Open(config.Type, config.ReadConnection)
	} else {
		Read = Write
	}
	Write.AutoMigrate(models...)
	return
}

func Close() {
	Write.Close()
	if len(config.ReadConnection) > 0 {
		Read.Close()
	}
}

func AddModel(m interface{}) {
	models = append(models, m)
}
