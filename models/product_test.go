package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExists(t *testing.T) {
	assert := assert.New(t)

	ok, err := (&Product{ID: 3}).Exists()
	assert.True(ok)
	assert.NoError(err)

	// test cache
	ok, err = (&Product{ID: 3}).Exists()
	assert.True(ok)
	assert.NoError(err)

	// WARNING: test cache after 5min skipped
}
