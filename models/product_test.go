package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExists(t *testing.T) {
	assert := assert.New(t)
	ok, err := ProductExists(3)

	assert.True(ok)
	assert.NoError(err)
}
