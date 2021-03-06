// Package with supporting functionality to run the microservice
package runtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// Function to test goodRelease()
func TestGoodRelease(t *testing.T) {
	assert := assert.New(t)
	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: ":memory:",
	})
	now := time.Now()
	good := models.Good{
		LockedAt:     &now,
		LockedSecret: "never used",
	}
	database.Write.Create(&good)
	count := GoodRelease(time.Duration(3) * time.Second)
	assert.Equal(int64(0), count, "no locked in timeout")

	older := now.Add(-time.Duration(10) * time.Minute)
	good.LockedAt = &older
	database.Write.Save(&good)
	count = GoodRelease(time.Duration(3) * time.Second)
	assert.Equal(int64(1), count, "unlock after timeout")

	grw := NewGoodReleaseWorker(models.GoodReleaseConfig{
		Every: models.Duration{Duration: time.Duration(3) * time.Millisecond},
		After: models.Duration{Duration: time.Duration(5) * time.Millisecond},
	})
	go grw.Start()
	time.Sleep(time.Duration(15) * time.Millisecond)
	grw.Close()

	database.Close()
}
