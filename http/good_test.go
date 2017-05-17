// Package that contains all api routes of this microservice
package http

import (
	"net/http"
	"testing"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/runtime"
	"github.com/genofire/hs_master-kss-monolith/test"
)

// Function to test addGood()
func TestAddGood(t *testing.T) {
	assertion, router := test.Init(t)

	BindAPI(router)
	runtime.PermissionURL = "http://localhost:8080/api-test/session/%s/%d/"
	session := test.NewSession(router)

	good := models.Good{
		ProductID: 3,
		Comment:   "blub",
	}

	_, w := session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusNonAuthoritativeInfo, w.StatusCode)

	session.Login()

	_, w = session.JSONRequest("POST", "/api/good/a", good)
	assertion.Equal(http.StatusNotAcceptable, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/good/7", good)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/good/1", true)
	assertion.Equal(http.StatusBadRequest, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusOK, w.StatusCode)

	database.Close()

	_, w = session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusInternalServerError, w.StatusCode)

	session.Logout()

	_, w = session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusForbidden, w.StatusCode)

	session.Login()
	runtime.CacheConfig.After = models.Duration{Duration: time.Duration(5) * time.Millisecond}
	test.CloseServer()
	time.Sleep(time.Duration(10) * time.Millisecond)
	runtime.HasPermission("testsessionkey", runtime.PermissionCreateGood)
	runtime.CleanCache()

	// Test gatewaytimeout on product exists
	_, w = session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusGatewayTimeout, w.StatusCode)

	time.Sleep(time.Duration(10) * time.Millisecond)
	runtime.CleanCache()

	// Test gatewaytimeout on permission exists
	_, w = session.JSONRequest("POST", "/api/good/1", good)
	assertion.Equal(http.StatusGatewayTimeout, w.StatusCode)

	test.Close()
}
