package http

import (
	"net/http"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

func status(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	var reviews []*models.Review
	var count int64
	database.Read.Find(&reviews).Count(&count)
	lib.Write(w, map[string]interface{}{"status": "running", "review_count": count})
	log.Info("done")
}
