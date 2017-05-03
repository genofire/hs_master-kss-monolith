// Package with supporting functionality to run the microservice
package runtime

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/worker"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// Configuration of the cache Worker
var CacheConfig models.CacheWorkerConfig

// Function to run the cache Worker
func CleanCache() {
	before := time.Now().Add(-CacheConfig.After.Duration)
	// Cache, if product exists
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

// Function to create a Worker and to clean the caches from other microservice
func NewCacheWorker() (w *worker.Worker) {
	return worker.NewWorker(CacheConfig.Every.Duration, CleanCache)
}
