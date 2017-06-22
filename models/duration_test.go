// Package with the mostly static content (models) of this microservice
package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Function to test UnmarshalTOML()
func TestDuration(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    interface{}
		err      string
		duration time.Duration
	}{
		{3, "invalid duration: \"%!s(int=3)\"", 0},
		{"", "invalid duration: \"\"", 0},
		{"am", "unable to parse duration am: strconv.Atoi: parsing \"a\": invalid syntax", 0},
		{"1x", "invalid duration unit: x", 0},
		{"1s", "", time.Second},
		{"73s", "", time.Second * 73},
		{"1m", "", time.Minute},
		{"73m", "", time.Minute * 73},
		{"1h", "", time.Hour},
		{"43h", "", time.Hour * 43},
		{"1d", "", time.Hour * 24},
		{"8d", "", time.Hour * 24 * 8},
		{"1w", "", time.Hour * 24 * 7},
		{"52w", "", time.Hour * 24 * 7 * 52},
		{"1y", "", time.Hour * 24 * 365},
		{"3y", "", time.Hour * 24 * 365 * 3},
	}

	for _, test := range tests {

		d := Duration{}
		err := d.UnmarshalTOML(test.input)
		duration := d.Duration

		if test.err == "" {
			assert.NoError(err)
			assert.Equal(test.duration, duration)
		} else {
			assert.EqualError(err, test.err)
		}
	}
}
