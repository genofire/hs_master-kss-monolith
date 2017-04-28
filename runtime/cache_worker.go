package runtime

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/worker"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// config of the cache worker
var CacheConfig models.CacheWorkerConfig

// command which is runned in the cache worker
func CleanCache() {
	before := time.Now().Add(-CacheConfig.After.Duration)
	// Cache if product exists
	for index, cache := range productExistCache {
		if before.After(cache.LastCheck) {
			delete(productExistCache, index)
		}
	}
	// Cache for permissions
	for index, cache := range permissionCache {
		if before.After(cache.LastCheck) {
			delete(permissionCache, index)
		}
	}
}

// create a worker to clean the caches which stored from other microservice
func NewCacheWorker() (w *worker.Worker) {
	return worker.NewWorker(CacheConfig.Every.Duration, CleanCache)
}
