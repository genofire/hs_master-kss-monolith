// Package that contains all api routes of this microservice
package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to test tempPercent() and tempProcessRadius()
func TestTempFuncs(t *testing.T) {
	assert := assert.New(t)
	resultInt := tempPercent(3, 9)
	assert.Equal(33, resultInt)

	// TODO is there a other way to calc this?
	resultFloat := tempProcessRadius(3, 9, 0)
	assert.Equal(float64(0), resultFloat)

	resultFloat = tempProcessRadius(12, 9, 10)
	assert.Equal(float64(0), resultFloat)
}
