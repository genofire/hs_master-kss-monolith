package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// TODO DRAFT for a rest request to a other microservice
const ProductURL = "https://google.com/?q=%d"

type boolMicroServiceCache struct {
	LastCheck time.Time
	Value     bool
}

var productExistCache map[int64]boolMicroServiceCache

func init() {
	productExistCache = make(map[int64]boolMicroServiceCache)
}

func (p *Product) Exists() (bool, error) {
	if cache, ok := productExistCache[p.ID]; ok {
		// cache for 5min
		before := time.Now().Add(-time.Minute * 5)
		if !cache.LastCheck.Before(before) {
			return cache.Value, nil
		}
	}

	url := fmt.Sprintf(ProductURL, p.ID)
	log.Log.WithField("url", url).Info("exists product?")
	res, err := http.Get(url)

	productExistCache[p.ID] = boolMicroServiceCache{
		LastCheck: time.Now(),
		Value:     (res.StatusCode == http.StatusOK),
	}
	return productExistCache[p.ID].Value, err
}
