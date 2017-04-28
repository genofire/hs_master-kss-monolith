package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

var (
	Write  *gorm.DB
	Read   *gorm.DB
	config *Config
	runtime []interface{}
)

type Config struct {
	Type           string
	Connection     string
	ReadConnection string
	Logging        bool
}

func Open(c Config) (err error) {
	writeLog := log.Log.WithField("db", "write")
	config = &c
	Write, err = gorm.Open(config.Type, config.Connection)
	if err != nil {
		return
	}
	Write.SingularTable(true)
	Write.LogMode(c.Logging)
	Write.SetLogger(writeLog)
	Write.Callback().Create().Remove("gorm:update_time_stamp")
	Write.Callback().Update().Remove("gorm:update_time_stamp")
	if len(config.ReadConnection) > 0 {
		readLog := log.Log.WithField("db", "read")
		Read, err = gorm.Open(config.Type, config.ReadConnection)
		if err != nil {
			return
		}
		Read.SingularTable(true)
		Read.LogMode(c.Logging)
		Read.SetLogger(readLog)
		Read.Callback().Create().Remove("gorm:update_time_stamp")
		Read.Callback().Update().Remove("gorm:update_time_stamp")
	} else {
		Read = Write
	}
	Write.AutoMigrate(runtime...)
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
	runtime = append(runtime, m)
}