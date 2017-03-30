package http

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"

	"github.com/genofire/hs_master-kss-monolith/test"
)

func TestReview(t *testing.T) {
	assertion, router := test.Init(t)

	BindAPI(router)
	session := test.NewSession(router)

	result, w := session.JSONRequest("GET", "/api/review/a", nil)
	assertion.Equal(http.StatusNotAcceptable, w.StatusCode)

	result, w = session.JSONRequest("GET", "/api/review/1", nil)
	assertion.Equal(http.StatusNotFound, w.StatusCode)

	database.Write.Create(&models.Review{
		ProductID:   3,
		FirstName:   "Max",
		LastName:    "Mustmann",
		RatingStars: 3,
		Text:        "blub",
	})

	result, w = session.JSONRequest("GET", "/api/review/3", nil)
	assertion.Equal(http.StatusOK, w.StatusCode)
	assertion.Len(result, 1)

	test.Close()
}
