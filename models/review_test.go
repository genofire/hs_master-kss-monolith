package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayName(t *testing.T) {
	assert := assert.New(t)

	r := Review{}
	assert.Equal("", r.FirstName, "wrong firstname")
	assert.Equal("", r.LastName, "wrong lastname")
	assert.Equal("Anonymous", r.DisplayName(), "No name")

	r.FirstName = "Max"
	assert.Equal("Max", r.FirstName, "wrong firstname")
	assert.Equal("", r.LastName, "wrong lastname")
	assert.Equal("Max", r.DisplayName(), "Only Firstname")

	r.LastName = "Mustermann"
	assert.Equal("Max", r.FirstName, "wrong firstname")
	assert.Equal("Mustermann", r.LastName, "wrong lastname")
	assert.Equal("Max M.", r.DisplayName(), "Shorted Name")

	r.FirstName = ""
	assert.Equal("", r.FirstName, "wrong firstname")
	assert.Equal("Mustermann", r.LastName, "wrong lastname")
	assert.Equal("Anonymous", r.DisplayName(), "displayname: no firstname")
}
