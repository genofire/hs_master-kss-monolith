package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestModel struct {
	ID    int64
	Value string `gorm:"type:varchar(255);column:value" json:"value"`
}

func TestOpenNoDB(t *testing.T) {
	assert := assert.New(t)

	c := Config{}

	err := Open(c)
	assert.Error(err, "error")
}
func TestOpenOneDB(t *testing.T) {
	assert := assert.New(t)
	AddModel(&TestModel{})

	c := Config{
		Type:       "sqlite3",
		Connection: "file:database?mode=memory",
	}
	var count int64

	err := Open(c)
	assert.NoError(err, "no error")
	Write.LogMode(true)
	Read.LogMode(true)

	Write.Create(&TestModel{Value: "first"})
	Write.Create(&TestModel{Value: "secound"})

	var list []*TestModel
	Read.Find(&list).Count(&count)
	assert.Equal(int64(2), count, "not enought entries")
	Close()
}

func TestOpenTwoDB(t *testing.T) {
	assert := assert.New(t)
	AddModel(&TestModel{})

	c := Config{
		Type:           "sqlite3",
		Connection:     "file:database?mode=memory",
		ReadConnection: "file:database2?mode=memory",
	}
	var count int64

	err := Open(c)
	assert.NoError(err, "no error")
	Write.LogMode(true)
	Read.LogMode(true)

	Write.Create(&TestModel{Value: "first"})
	Write.Create(&TestModel{Value: "secound"})

	var list []*TestModel
	Write.Find(&list).Count(&count)
	assert.Equal(int64(2), count, "not enought entries")

	result := Read.Find(&list)
	assert.Error(result.Error, "error, because it is the wrong database")
	Close()

}
