// Package with supporting functionality to run the microservice
package runtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// Function to test fouledDelete()
func TestFouledDelete(t *testing.T) {
	assert := assert.New(t)
	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: ":memory:",
	})

	now := time.Now().Add(-time.Hour * 48)
	good := models.Good{
		FouledAt: &now,
	}
	database.Write.Create(&good)

	good = models.Good{
		FouledAt: &now,
	}
	database.Write.Create(&good)

	count := GoodFouled()
	assert.Equal(2, count, "not fouled")

	good = models.Good{
		FouledAt: &now,
	}
	database.Write.Create(&good)

	now = time.Now().Add(time.Hour * 48)
	good = models.Good{
		FouledAt: &now,
	}
	database.Write.Create(&good)

	count = GoodFouled()
	assert.Equal(1, count, "fouled")

	database.Close()
}
