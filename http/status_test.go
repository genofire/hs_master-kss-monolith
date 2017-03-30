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

	r, w := session.JSONRequest("GET", "/api/status", nil)
	result := r.(map[string]interface{})
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal("running", result["status"])
	assertion.Equal(float64(0), result["review_count"])

}
