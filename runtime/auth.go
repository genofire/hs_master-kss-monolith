package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

// url to the microservice which manage the permissions
var PermissionURL string

// type of permission
type Permission int

// some permission (see Permission)
const (
	// has the user the permission to create need goods of a product
	// e.g. if a good received and now availablity to sell
	PermissionCreateGood = 1
	// has the user the permission to delete need goods of a product
	// e.g. if a good become rancid and has to remove from stock
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

	value := false
	if err == nil {
		value = (res.StatusCode == http.StatusOK)
	}

	c.permissions[p] = boolMicroServiceCache{
		LastCheck: c.LastCheck,
		Value:     value,
	}
	return c.permissions[p].Value, err
}

var permissionCache map[string]*permissionMicroServiceCache

func init() {
	permissionCache = make(map[string]*permissionMicroServiceCache)
}

// check if the client with the session string has a permissions (see Permission)
func HasPermission(session string, p int) (bool, error) {
	_, ok := permissionCache[session]
	if !ok {
		permissionCache[session] = &permissionMicroServiceCache{
			LastCheck:   time.Now(),
			session:     session,
			permissions: make(map[Permission]boolMicroServiceCache),
		}
	}
	return permissionCache[session].HasPermission(Permission(p))
}
