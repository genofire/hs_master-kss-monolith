// Package that contains all api routes of this microservice
package http

import (
	goji "goji.io"
	"goji.io/pat"

	"github.com/genofire/hs_master-kss-monolith/lib/http"
	"github.com/genofire/hs_master-kss-monolith/runtime"
)

// Function to bind all api routes to the webserver
func BindAPI(router *goji.Mux) {
	router.HandleFunc(pat.Get("/api/status"), status)
	router.HandleFunc(pat.Get("/api/good/:productid"), listGoods)
	router.HandleFunc(pat.Get("/api/good/availablity/:productid"), getGoodAvailablity)
	router.HandleFunc(pat.Post("/api/good/:productid"), http.PermissionHandler(addGood, runtime.HasPermission, runtime.PermissionCreateGood))
}
