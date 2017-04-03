package http

import (
	goji "goji.io"
	"goji.io/pat"
)

func BindAPI(router *goji.Mux) {
	router.HandleFunc(pat.Get("/api/status"), status)
	router.HandleFunc(pat.Get("/api/good/:productid"), listGoods)
}
