package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

var PermissionURL string

type Permission int

const (
	PermissionCreateGood = 1
	PermissionDeleteGood = 2
)

type permissionMicroServiceCache struct {
	LastCheck   time.Time
	session     string
	permissions map[Permission]boolMicroServiceCache
}

func (c *permissionMicroServiceCache) HasPermission(p Permission) (bool, error) {
	c.LastCheck = time.Now()
	if cache, ok := c.permissions[p]; ok {
		before := time.Now().Add(-CacheConfig.After.Duration)
		if before.After(cache.LastCheck) {
			return cache.Value, nil
		}
	}

	url := fmt.Sprintf(PermissionURL, c.session, p)
	log.Log.WithField("url", url).Info("has permission?")
	res, err := http.Get(url)

	c.permissions[p] = boolMicroServiceCache{
		LastCheck: c.LastCheck,
		Value:     (res.StatusCode == http.StatusOK),
	}
	return c.permissions[p].Value, err
}

var permissionCache map[string]*permissionMicroServiceCache

func init() {
	permissionCache = make(map[string]*permissionMicroServiceCache)
}

func HasPermission(session string, p Permission) (bool, error) {
	_, ok := permissionCache[session]
	if !ok {
		permissionCache[session] = &permissionMicroServiceCache{
			LastCheck:   time.Now(),
			session:     session,
			permissions: make(map[Permission]boolMicroServiceCache),
		}
	}
	return permissionCache[session].HasPermission(p)
}