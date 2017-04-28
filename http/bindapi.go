// all api routes of this microservice
package http

import (
	goji "goji.io"
	"goji.io/pat"
)

// bind all API routes to webserver
func BindAPI(router *goji.Mux) {
	router.HandleFunc(pat.Get("/api/status"), status)
	router.HandleFunc(pat.Get("/api/good/:productid"), listGoods)
	router.HandleFunc(pat.Get("/api/good/availablity/:productid"), getGoodAvailablity)
}
