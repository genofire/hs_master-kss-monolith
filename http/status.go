package http

import (
	"net/http"

	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	lib.Write(w, "running")
	log.Info("show status")
}
