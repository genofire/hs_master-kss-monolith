package models

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/worker"
)

var CacheConfig CacheWorkerConfig

type CacheWorkerConfig struct {
	Every Duration
	After Duration
}

func NewCacheWorker() (w *worker.Worker) {
	return worker.NewWorker(CacheConfig.Every.Duration, func() {
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
	})
}
