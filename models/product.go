package models

import (
	"net/http"
	"time"
)

type boolMicroServiceCache struct {
	LastCheck time.Time
	Value     bool
}

var productExistCache map[int64]boolMicroServiceCache

func init() {
	productExistCache = make(map[int64]boolMicroServiceCache)
}

func ProductExists(id int64) (bool, error) {
	if cache, ok := productExistCache[id]; ok {
		// cache for 5min
		before := time.Now().Add(-time.Minute * 5)
		if !cache.LastCheck.Before(before) {
			return cache.Value, nil
		}
	}

	// TODO DRAFT for a rest request to a other microservice
	res, err := http.Get("http://golang.org")

	productExistCache[id] = boolMicroServiceCache{
		LastCheck: time.Now(),
		Value:     (res.StatusCode == http.StatusOK),
	}
	return productExistCache[id].Value, err
}
