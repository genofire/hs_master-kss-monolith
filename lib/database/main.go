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
	Logging bool
}

func Open(c Config) (err error) {
	config = &c
	Write, err = gorm.Open(config.Type, config.Connection)
	Write.SingularTable(true)
	Write.LogMode(c.Logging)
	Write.Callback().Create().Remove("gorm:update_time_stamp")
	Write.Callback().Update().Remove("gorm:update_time_stamp")
	if err != nil {
		return
	}
	if len(config.ReadConnection) > 0 {
		Read, err = gorm.Open(config.Type, config.ReadConnection)
		Read.SingularTable(true)
		Read.LogMode(c.Logging)
		Read.Callback().Create().Remove("gorm:update_time_stamp")
		Read.Callback().Update().Remove("gorm:update_time_stamp")
	} else {
		Read = Write
	}
	Write.AutoMigrate(models...)
	return
}

func Close() {
	Write.Close()
	Write = nil

	if len(config.ReadConnection) > 0 {
		Read.Close()
	}
	Read = nil
}

func AddModel(m interface{}) {
	models = append(models, m)
}
