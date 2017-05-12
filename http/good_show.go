// Package that contains all api routes of this microservice
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
	"github.com/genofire/hs_master-kss-monolith/runtime"
	"time"
)

// Function to list all goods
func listGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "the productid is false", http.StatusNotAcceptable)
		return
	}
	log = log.WithField("productid", id)
	var list []*models.Good
	result := database.Read.Where("product_id = ?", id).Find(&list)
	if result.RowsAffected == 0 {
		log.Warn("no goods found")
		http.Error(w, "no goods found for this product", http.StatusNotFound)
		return
	}
	lib.Write(w, list)
	log.Info("done")
}

// Function that counts als available goods for one product
func getGoodAvailablityCount(w http.ResponseWriter, r *http.Request) (int, *logrus.Entry) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("wrong productid format")
		http.Error(w, "the product id has a false format", http.StatusNotAcceptable)
		return -1, log
	}
	log = log.WithField("productid", id)
	ok, err := runtime.ProductExists(id)
	if err != nil {
		log.Warn("product could not verified on the microservice")
		http.Error(w, "the product could not be verified", http.StatusGatewayTimeout)
		return -1, log
	}
	if !ok {
		log.Warn("product does not exists anymore")
		http.Error(w, "the product does not exists anymore", http.StatusNotFound)
		return -1, log
	}
	var count float64
	(&models.Good{}).FilterAvailable(database.Read.Where("product_id = ?", id)).Count(&count)
	return int(count), log
}

// Function that returns the availability of a good
func getGoodAvailability(w http.ResponseWriter, r *http.Request) {
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



// Function that returns the freshness of a good
func getGoodFreshness(w http.ResponseWriter, r *http.Request){
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "goodid"), 10, 64)
	if err != nil {
		log.Warn("wrong goodid format")
		http.Error(w, "the good id has a false format", http.StatusNotAcceptable)
		return
	}
	log = log.WithField("goodid", id)

	var good models.Good
	database.Read.Where("id = ?", id).First(&good)
	if good.ProductID == 0 {
		log.Warn("good not found")
		http.Error(w, "the good was not found in database", http.StatusNotFound)
		return
	}

	fresh := good.FouledAt.Before(time.Now().Add(-time.Duration(3) * time.Hour * 24))

	log = log.WithField("type", r.Header.Get("Content-Type"))
	switch r.Header.Get("Content-Type") {
	case "application/json":
		lib.Write(w, fresh)
	default:
		getGoodFreshnessSVG(w, fresh)
	}
	log.Info("done")
}