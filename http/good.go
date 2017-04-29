package http

import (
	"net/http"
	"strconv"

	"goji.io/pat"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/runtime"
)

func addGood(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "wrong productid", http.StatusNotAcceptable)
		return
	}
	log = log.WithField("productid", id)
	ok, err := runtime.ProductExists(id)
	if err != nil {
		log.Warn(err.Error())
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	if !ok {
		log.Warn("wrong product not found")
		http.Error(w, "wrong product not found", http.StatusNotFound)
		return
	}

	var obj models.Good
	lib.Read(r, &obj)

	obj.ProductID = id

	db := database.Write.Create(&obj)
	if db.Error != nil {
		log.Error("database could not write", db.Error)
		http.Error(w, "was not possible to write", http.StatusInternalServerError)
	}
	lib.Write(w, &obj)

	log.Info("done")
}
