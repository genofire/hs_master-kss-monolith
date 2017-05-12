// Package with supporting functionality to run the microservice
package runtime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/log"
	"sync"
)

// URL to the microservice which manages permissions
var PermissionURL string

// Type of permission
type Permission int

// Some permissions (the real permissions need to come from the permission microservice)
const (
	// permission to add goods to the stock
	// e.g. if a good is received and now available to sell
	PermissionCreateGood = 1

	// permission to delete goods from the stock
	// e.g. if a good become rancid and has to be removed
	PermissionDeleteGood = 2
)

// Struct that holds the information for a permission cache
type permissionMicroServiceCache struct {
	LastCheck   time.Time
	session     string
	permissions map[Permission]boolMicroServiceCache
	sync.Mutex
}


// Function to check, if a user has a permission
func (c *permissionMicroServiceCache) HasPermission(p Permission) (bool, error) {
	c.LastCheck = time.Now()
	c.Lock()
	defer c.Unlock()
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

// Cache for permissions
var permissionCache map[string]*permissionMicroServiceCache
var permissionMutex sync.Mutex
// Function to initialize the permission cache
func init() {
	permissionCache = make(map[string]*permissionMicroServiceCache)
}

// Function to check, if the current session has any permissions
func HasPermission(session string, p int) (bool, error) {
	permissionMutex.Lock()
	defer permissionMutex.Unlock()
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
