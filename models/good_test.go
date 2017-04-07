package models

import (
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/stretchr/testify/assert"
)

func TestGood(t *testing.T) {
	assert := assert.New(t)

	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: ":memory:",
	})

	good := &Good{}
	assert.False(good.IsLock())

	good.Lock("blub_secret")
	assert.True(good.IsLock())

	err := good.Unlock("secret")
	assert.Error(err)
	assert.True(good.IsLock())

	good.Unlock("blub_secret")
	assert.False(good.IsLock())

	assert.NotNil(good.FilterAvailable(database.Read))

}
