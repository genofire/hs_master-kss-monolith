package http

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"

	"github.com/genofire/hs_master-kss-monolith/test"
)

func TestGood(t *testing.T) {
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
