package http

import (
	"net/http"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

func listReview(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	var list []*models.Review
	database.Read.Find(&list)
	lib.Write(w, list)
	log.Info("done")
}
