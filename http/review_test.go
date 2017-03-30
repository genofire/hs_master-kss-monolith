package http

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"

	"github.com/genofire/hs_master-kss-monolith/test"
)

func TestReview(t *testing.T) {
	database.Open(database.Config{
		Type:       "sqlite3",
		Connection: ":memory:",
	})
	database.Write.LogMode(true)
	assertion, router := test.Init(t)
	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/reviews", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal([]interface{}{}, result)
}
