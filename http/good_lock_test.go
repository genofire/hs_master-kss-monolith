// Package that contains all api routes of this microservice
package http

import (
	"net/http"
	"testing"
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/test"
)

// Function to test lockGoods()
func TestLockGoods(t *testing.T) {
	assertion, router := test.Init(t)
	good := &models.Good{
		ProductID: 3,
		Comment:   "blabla",
	}
	database.Write.Create(good)

	BindAPI(router)
	session := test.NewSession(router)

	_, w := session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 3, Count: 1}})
	assertion.Equal(http.StatusBadRequest, w.StatusCode)

	session.Header["secret"] = "hiddenLockTest"

	_, w = session.JSONRequest("POST", "/api/goods/locking", 13)
	assertion.Equal(http.StatusBadRequest, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/goods/locking", nil)
	assertion.Equal(http.StatusBadRequest, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 0, Count: 2}})
	assertion.Equal(http.StatusBadRequest, w.StatusCode)

	_, w = session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 3, Count: 2}})
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	result, w := session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 3, Count: 1}})
	assertion.Equal(http.StatusOK, w.StatusCode)
	resultMap := result.(map[string]interface{})
	count := resultMap["count"]
	assertion.Equal(float64(1), count)

	_, w = session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 3, Count: 1}})
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	database.Close()

	assertion.Panics(func() {
		_, w = session.JSONRequest("POST", "/api/goods/locking", []interface{}{LockGood{ProductID: 3, Count: 1}})
		assertion.Equal(http.StatusInternalServerError, w.StatusCode)
	})

	test.Close()
}

// Function to test releaseGoods()
func TestReleaseGoods(t *testing.T) {
	now := time.Now()
	assertion, router := test.Init(t)

	database.Write.Create(&models.Good{
		ProductID:    3,
		Comment:      "blub",
		LockedAt:     &now,
		LockedSecret: "hidden",
	})

	BindAPI(router)
	session := test.NewSession(router)

	session.Header["secret"] = "a"
	result, w := session.JSONRequest("DELETE", "/api/goods/locking", nil)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	session.Header["secret"] = "hidden"
	result, w = session.JSONRequest("DELETE", "/api/goods/locking", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	resultMap := result.(map[string]interface{})
	count := resultMap["count"]
	assertion.Equal(float64(1), count)

	database.Close()
	result, w = session.JSONRequest("DELETE", "/api/goods/locking", nil)
	assertion.Equal(http.StatusInternalServerError, w.StatusCode)

	test.Close()

}
