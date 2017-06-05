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
