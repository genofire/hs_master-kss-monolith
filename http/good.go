// Package that contains all api routes of this microservice
package http

import (
	"net/http"
	"strconv"
	"time"

	"goji.io/pat"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/runtime"
	"github.com/jinzhu/gorm"
)

// Function to add goods to the stock
func addGood(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)

	countStr := r.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 0
	}

	id, err := strconv.ParseInt(pat.Param(r, "productid"), 10, 64)
	if err != nil {
		log.Warn("false product id format")
		http.Error(w, "the product id has a false format", http.StatusNotAcceptable)
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
		log.Warn("product not found")
		http.Error(w, "the product was not found", http.StatusNotFound)
		return
	}

	var obj models.Good
	err = lib.Read(r, &obj)
	if err != nil {
		log.Warn(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj.ProductID = id

	var db *gorm.DB
	if count > 0 {
		for i := 0; i < count; i++ {
			db = database.Write.Create(&obj)
			obj.ID = 0
		}
	} else {
		db = database.Write.Create(&obj)
	}

	if db.Error != nil {
		log.Error("database unable to write", db.Error)
		http.Error(w, "the product could not be written into the database", http.StatusInternalServerError)
	}
	lib.Write(w, &obj)

	log.Info("done")
}

// Function that returns the freshness of a good
func delGood(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	id, err := strconv.ParseInt(pat.Param(r, "goodid"), 10, 64)
	if err != nil {
		log.Warn("wrong good id format")
		http.Error(w, "the good id has a false format", http.StatusNotAcceptable)
		return
	}
	log = log.WithField("goodid", id)

	now := time.Now()
	var good models.Good
	good.ID = id
	db := good.FilterAvailable(database.Read).First(&good)
	if db.RecordNotFound() {
		log.Warnf("could not find good: %s", db.Error)
		http.Error(w, "the good was not found", http.StatusNotFound)
		return
	}
	good.ManuelleDelete = true
	good.DeletedAt = &now

	db = database.Write.Save(&good)
	if db.Error != nil {
		log.Warnf("could not delete good: %s", db.Error)
		http.Error(w, "the good could not be deleted", http.StatusInternalServerError)
		return
	}
	log.Info("done")
}
