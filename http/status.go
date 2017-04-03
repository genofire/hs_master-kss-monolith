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
	var good []*models.Good
	var count int64
	var avg int64
	database.Read.Find(&good).Count(&count) //.Avg(&avg)
	lib.Write(w, map[string]interface{}{
		"status": "running",
		"good": map[string]interface{}{
			"count": count,
			"avg":   avg,
		},
	})
	log.Info("done")
}
