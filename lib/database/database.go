// Package database provides the
// functionality to open, close and use a database
package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// Database connection for writing
var Write *gorm.DB

// Database connection for reading
var Read *gorm.DB

var (
	config  *Config
	runtime []interface{}
)

// configuration for the database connection
type Config struct {
	// type of database: current support sqlite and postgres
	// (by request other could be enabled)
	Type string
	// connection configuration
	Connection string
	// maybe create another connection just for reading
	ReadConnection string
	// enable logging the generated sql string
	Logging bool
}

// Function to open a database and set the given configuration
// Input: the configuration data c
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

// Function to safely close the database
func Close() {
	Write.Close()
	Write = nil
	if len(config.ReadConnection) > 0 {
		Read.Close()
	}
	Read = nil
}

// Function to add a model to the runtime
// Input: interface m
func AddModel(m interface{}) {
	runtime = append(runtime, m)
}
