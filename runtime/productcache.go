package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// url to the microservice which manage the products
var ProductURL string

type boolMicroServiceCache struct {
	LastCheck time.Time
	Value     bool
}

var productExistCache map[int64]boolMicroServiceCache

func init() {
	productExistCache = make(map[int64]boolMicroServiceCache)
}

// check on the other microservice if the product exists
func ProductExists(id int64) (bool, error) {
	if cache, ok := productExistCache[id]; ok {
		return cache.Value, nil
	}

	url := fmt.Sprintf(ProductURL, id)
	log.Log.WithField("url", url).Info("exists product?")
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
