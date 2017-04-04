package http

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/test"
)

func TestStatus(t *testing.T) {
	assertion, router := test.Init(t)

	BindAPI(router)
	session := test.NewSession(router)

	database.Write.Create(&models.Good{
		ProductID: 3,
		Position:  "regal 1",
	})
	database.Write.Create(&models.Good{
		ProductID: 3,
		Position:  "regal 2",
	})
	database.Write.Create(&models.Good{
		ProductID: 1,
		Position:  "regal 10",
	})

	r, w := session.JSONRequest("GET", "/api/status", nil)
	result := r.(map[string]interface{})
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Equal("running", result["status"])

	db := result["database"].(map[string]interface{})
	assertion.Equal(true, db["read"])
	assertion.Equal(true, db["write"])

	good := result["good"].(map[string]interface{})
	assertion.Equal(float64(3), good["count"])
	assertion.Equal(float64(1.5), good["avg"])

	test.Close()
}
