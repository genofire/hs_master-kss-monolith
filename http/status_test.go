package http

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/test"
)

func TestStatus(t *testing.T) {
	assertion, router := test.Init(t)
	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/status", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal("running", result)

}
