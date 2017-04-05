package models

import (
	"testing"
	"time"
)

func TestCacheWorker(t *testing.T) {

	productExistCache[2] = boolMicroServiceCache{LastCheck: time.Now(), Value: true}
	permissionCache["blub"] = &permissionMicroServiceCache{
		LastCheck:   time.Now(),
		session:     "blub",
		permissions: make(map[Permission]boolMicroServiceCache),
	}
	CacheConfig = CacheWorkerConfig{
		Every: Duration{Duration: time.Duration(3) * time.Millisecond},
		After: Duration{Duration: time.Duration(5) * time.Millisecond},
	}
	cw := NewCacheWorker()
	go cw.Start()
	time.Sleep(time.Duration(15) * time.Millisecond)
	cw.Close()
}
