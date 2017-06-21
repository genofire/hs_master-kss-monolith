// Package with supporting functionality to run the microservice
package runtime

import (
	"testing"
	"time"

	"github.com/genofire/hs_master-kss-monolith/models"
)

// Function to test the cacheWorker
func TestCacheWorker(t *testing.T) {

	productExistCache[2] = boolMicroServiceCache{LastCheck: time.Now(), Value: true}
	permissionCache["blub"] = &permissionMicroServiceCache{
		LastCheck:   time.Now(),
		session:     "blub",
		permissions: make(map[Permission]boolMicroServiceCache),
	}
	CacheConfig = models.CacheWorkerConfig{
		Every: models.Duration{Duration: time.Duration(3) * time.Millisecond},
		After: models.Duration{Duration: time.Duration(5) * time.Millisecond},
	}
	cw := NewCacheWorker()
	go cw.Start()
	time.Sleep(time.Duration(15) * time.Millisecond)
	cw.Close()
}
