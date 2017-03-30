package http

import (
	goji "goji.io"
	"goji.io/pat"
)

func BindAPI(router *goji.Mux) {
	router.HandleFunc(pat.Get("/api/status"), statusHandler)
}