package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGood(t *testing.T) {
	assert := assert.New(t)

	good := &Good{}
	assert.False(good.IsLock())

	good.Lock("blub_secret")
	assert.True(good.IsLock())

	err := good.Unlock("secret")
	assert.Error(err)
	assert.True(good.IsLock())

	good.Unlock("blub_secret")
	assert.False(good.IsLock())
}
