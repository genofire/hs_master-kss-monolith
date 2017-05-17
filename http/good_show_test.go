// Package that contains all api routes of this microservice
package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/runtime"
	"github.com/genofire/hs_master-kss-monolith/test"
)

// Function to test listGoods()
func TestListGood(t *testing.T) {
	assertion, router := test.Init(t)

	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/good/a", nil)
	assertion.Equal(http.StatusNotAcceptable, w.StatusCode)

	result, w = session.JSONRequest("GET", "/api/good/1", nil)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	database.Write.Create(&models.Good{
		ProductID: 3,
		Comment:   "blub",
	})

	result, w = session.JSONRequest("GET", "/api/good/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Len(result, 1)

	test.Close()
}

// Function to getGoodAvailability() and getGoodAvailabilityCount()
func TestGetGoodAvailable(t *testing.T) {
	now := time.Now()
	assertion, router := test.Init(t)

	runtime.ProductURL = "http://localhost:8080/api-test/product/%d/"

	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/good/availablity/a", nil)
	assertion.Equal(http.StatusNotAcceptable, w.StatusCode)

	result, w = session.JSONRequest("GET", "/api/good/availablity/1", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal(float64(0), result)

	database.Write.Create(&models.Good{
		ProductID:    3,
		Comment:      "blub",
		LockedAt:     &now,
		LockedSecret: "hidden",
	})
	database.Write.Create(&models.Good{
		ProductID: 3,
		Comment:   "blub",
	})
	database.Write.Create(&models.Good{
		ProductID: 3,
		Comment:   "blub",
	})

	result, w = session.JSONRequest("GET", "/api/good/availablity/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal(float64(2), result)

	req, _ := http.NewRequest("GET", "/api/good/availablity/3", nil)
	req.Header.Set("Content-Type", "image/svg+xml")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assertion.Equal(http.StatusOK, w.StatusCode)

	database.Write.Create(&models.Good{
		ProductID: 4,
		Comment:   "blub",
	})

	result, w = session.JSONRequest("GET", "/api/good/availablity/7", nil)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	test.CloseServer()
	runtime.CacheConfig.After = models.Duration{Duration: time.Duration(5) * time.Millisecond}
	time.Sleep(time.Duration(10) * time.Millisecond)
	runtime.CleanCache()

	result, w = session.JSONRequest("GET", "/api/good/availablity/3", nil)
	assertion.Equal(http.StatusGatewayTimeout, w.StatusCode)
	test.Close()

}

// Function to getGoodFreshness()
func TestGetGoodFreshness(t *testing.T) {
	now := time.Now().Add(36 * time.Hour)
	assertion, router := test.Init(t)

	runtime.ProductURL = "http://localhost:8080/api-test/product/%d/"

	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/good/freshness/a", nil)
	assertion.Equal(http.StatusNotAcceptable, w.StatusCode)

	database.Write.Create(&models.Good{
		ID:        3,
		ProductID: -2,
	})

	result, w = session.JSONRequest("GET", "/api/good/freshness/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal(false, result)

	database.Write.Save(&models.Good{
		ID:        3,
		ProductID: -2,
		FouledAt:  &now,
	})

	result, w = session.JSONRequest("GET", "/api/good/freshness/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal(true, result)

	fouled := now.Add(-72 * time.Hour)
	database.Write.Save(&models.Good{
		ID:        3,
		ProductID: -2,
		FouledAt:  &fouled,
	})
	result, w = session.JSONRequest("GET", "/api/good/freshness/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal(false, result)

	req, _ := http.NewRequest("GET", "/api/good/freshness/3", nil)
	req.Header.Set("Content-Type", "image/svg+xml")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assertion.Equal(http.StatusOK, w.StatusCode)

	result, w = session.JSONRequest("GET", "/api/good/freshness/7", nil)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	test.Close()

}
