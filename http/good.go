package http

import (
	"net/http"
	"strconv"

	logrus "github.com/Sirupsen/logrus"
	"goji.io/pat"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

func listGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "wrong productid", http.StatusNotAcceptable)
		return
	}
	log = log.WithField("productid", id)
	var list []*models.Good
	result := database.Read.Where("product_id = ?", id).Find(&list)
	if result.RowsAffected == 0 {
		log.Warn("no goods found")
		http.Error(w, "no goods found", http.StatusNotFound)
		return
	}
	lib.Write(w, list)
	log.Info("done")
}

func getGoodAvailablityCount(w http.ResponseWriter, r *http.Request) (int, *logrus.Entry) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "wrong productid", http.StatusNotAcceptable)
		return -1, log
	}
	log = log.WithField("productid", id)
	product := models.Product{ID: id}
	ok, err := product.Exists()
	if err != nil {
		log.Warn("product could not verified on other microservice")
		http.Error(w, "product could not verified on other microservice", http.StatusGatewayTimeout)
		return -1, log
	}
	if !ok {
		log.Warn("product did not exists anymore")
		http.Error(w, "product did not exists anymore", http.StatusNotFound)
		return -1, log
	}
	var count float64
	(&models.Good{}).FilterAvailable(database.Read.Where("product_id = ?", product.ID)).Count(&count)
	return int(count), log
}
func getGoodAvailablity(w http.ResponseWriter, r *http.Request) {
	count, log := getGoodAvailablityCount(w, r)
	if count < 0 {
		return
	}
	log = log.WithField("type", r.Header.Get("Content-Type"))
	switch r.Header.Get("Content-Type") {
	case "application/json":
		lib.Write(w, count)
	default:
		getGoodAvailablitySVG(w, count)
	}
	log.Info("done")
}
