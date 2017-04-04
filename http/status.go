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
	var goods []*models.Good
	var count int64
	var avg float64
	database.Read.Find(&goods).Count(&count)
	database.Read.Raw("SELECT avg(g.gcount) as avg FROM (select count(*) as gcount FROM good g WHERE deleted_at is NULL GROUP BY g.product_id) g").Row().Scan(&avg)
	lib.Write(w, map[string]interface{}{
		"status": "running",
		"database": map[string]interface{}{
			"read":  database.Read.DB().Ping() == nil,
			"write": database.Write.DB().Ping() == nil,
		},
		"good": map[string]interface{}{
			"count": count,
			"avg":   avg,
		},
	})
	log.Info("done")
}
