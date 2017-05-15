// Package with supporting functionality to run the microservice
package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
	"sync"
)

// URL to the microservice which manages the products (product catalogue)
var ProductURL string

// Struct that holds the information on the microservice cache
type boolMicroServiceCache struct {
	LastCheck time.Time
	Value     bool
}

// Cache for existing products
var productExistCache map[int64]boolMicroServiceCache
var productMutex sync.Mutex

// Function to initialize the cache for existing products
func init() {
	productExistCache = make(map[int64]boolMicroServiceCache)
}

// Function to check on the other microservice (product catalogue) if the product exists
func ProductExists(id int64) (bool, error) {
	productMutex.Lock()
	defer productMutex.Unlock()
	if cache, ok := productExistCache[id]; ok {
		return cache.Value, nil
	}

	url := fmt.Sprintf(ProductURL, id)
	log.Log.WithField("url", url).Info("does the product exist?")
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	productExistCache[id] = boolMicroServiceCache{
		LastCheck: time.Now(),
		Value:     (res.StatusCode == http.StatusOK),
	}
	return productExistCache[id].Value, nil
}
