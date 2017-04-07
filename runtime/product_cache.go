package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// TODO DRAFT for a rest request to a other microservice
const ProductURL = "http://localhost:8080/api-test/product/%d/"

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
		return cache.Value, nil
	}

	url := fmt.Sprintf(ProductURL, p.ID)
	log.Log.WithField("url", url).Info("exists product?")
	res, err := http.Get(url)
	if err == nil {
		productExistCache[p.ID] = boolMicroServiceCache{
			LastCheck: time.Now(),
			Value:     (res.StatusCode == http.StatusOK),
		}
	}
	return productExistCache[p.ID].Value, err
}
