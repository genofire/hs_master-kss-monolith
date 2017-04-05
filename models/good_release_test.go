package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
)

func TestGoodRelease(t *testing.T) {
	assert := assert.New(t)
	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: ":memory:",
	})
	now := time.Now()
	good := Good{
		LockedAt:     &now,
		LockedSecret: "never used",
	}
	database.Write.Create(&good)
	count := goodRelease(time.Duration(3) * time.Second)
	assert.Equal(int64(0), count, "no locked in timeout")

	older := now.Add(-time.Duration(10) * time.Minute)
	good.LockedAt = &older
	database.Write.Save(&good)
	count = goodRelease(time.Duration(3) * time.Second)
	assert.Equal(int64(1), count, "unlock after timeout")

	grw := NewGoodReleaseWorker(GoodReleaseConfig{
		Every: Duration{Duration: time.Duration(3) * time.Millisecond},
		After: Duration{Duration: time.Duration(5) * time.Millisecond},
	})
	go grw.Start()
	time.Sleep(time.Duration(15) * time.Millisecond)
	grw.Close()
}
