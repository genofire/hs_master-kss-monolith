package http

import (
	"net/http"
	"strconv"

	"goji.io/pat"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

func listReview(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "wrong productid", http.StatusNotAcceptable)
		return
	}
	log.WithField("productid", id)
	var list []*models.Review
	result := database.Read.Where("product_id = ?", id).Find(&list)
	if result.RowsAffected == 0 {
		log.Warn("no reviews found")
		http.Error(w, "no reviews found", http.StatusNotFound)
		return
	}
	lib.Write(w, list)
	log.Info("done")
}
