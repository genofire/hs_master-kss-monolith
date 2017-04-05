package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)

	perm, err := HasPermission("session", PermissionCreateGood)
	assert.NoError(err)
	assert.True(perm)
}
