package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTempFuncs(t *testing.T) {
	assert := assert.New(t)
	resultInt := tempProcent(3, 9)
	assert.Equal(33, resultInt)

	// TODO is there a other way to calc this?
	resultFloat := tempProcessRadius(3, 9, 0)
	assert.Equal(float64(0), resultFloat)
}
