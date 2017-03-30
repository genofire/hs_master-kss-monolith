package http

import (
	"net/http"

	"github.com/genofire/hs_master-kss-monolith/lib"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log := lib.LogHTTP(r)
	lib.Write(w, "running")
	log.Info("show status")
}
