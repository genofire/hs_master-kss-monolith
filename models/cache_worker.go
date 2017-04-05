package models

import "time"

type CacheWorkerConfig struct {
	Every Duration
	After Duration
}

func NewCacheWorker(config CacheWorkerConfig) (w *Worker) {
	return NewWorker(config.Every.Duration, func() {
		// Cache if product exists
		for index, cache := range productExistCache {
			if cache.LastCheck.After(time.Now().Add(-config.After.Duration)) {
				delete(productExistCache, index)
			}
		}
		// Cache for permissions
		for index, cache := range permissionCache {
			if cache.LastCheck.After(time.Now().Add(-config.After.Duration)) {
				delete(permissionCache, index)
			}
		}
	})
}
