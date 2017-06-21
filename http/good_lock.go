// Package that contains all api routes of this microservice
package http

import (
	"net/http"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	lib "github.com/genofire/hs_master-kss-monolith/lib/http"
	logger "github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

type LockGood struct {
	ProductID int64 `json:"product_id"`
	Count     int   `json:"count"`
}

// Function to lock goods
func lockGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	secret := r.Header.Get("secret")

	if secret == "" {
		log.Warn("no secred for locking given")
		http.Error(w, "no secred for locking given", http.StatusBadRequest)
		return
	}

	log = log.WithField("lSecret", secret)

	var goods []*LockGood

	err := lib.Read(r, &goods)

	if err != nil {
		log.Warn(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(goods) <= 0 {
		log.Warn("tried to log nothing")
		http.Error(w, "tried to log nothing", http.StatusBadRequest)
		return
	}

	tx := database.Write.Begin()
	count := int64(0)

	for _, good := range goods {
		if good.ProductID <= 0 {
			log.Warn("tried to log nothing")
			tx.Rollback()
			http.Error(w, "tried to log nothing", http.StatusBadRequest)
			return
		}
		for i := 0; i < good.Count; i++ {
			g := &models.Good{ProductID: good.ProductID}
			db := g.FilterAvailable(tx).First(g)
			if db.RecordNotFound() {
				log.Warn("good not found")
				tx.Rollback()
				http.Error(w, "the good was not found in database", http.StatusNotFound)
				return
			}
			g.Lock(secret)

			db = tx.Save(g)

			if db.Error != nil || db.RowsAffected != 1 {
				http.Error(w, "the good was not found in database", http.StatusInternalServerError)
				tx.Rollback()
				log.Panic("there is more than one good locked: ", db.Error)
				return
			}
			count += 1
		}
	}

	lib.Write(w, map[string]int64{"count": count})
	tx.Commit()
	log.Info("done")

}

// Function to release locked goods
func releaseGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	secret := r.Header.Get("secret")
	log = log.WithField("lSecret", secret)

	db := database.Write.Model(&models.Good{}).Where(&models.Good{LockedSecret: secret}).Updates(map[string]interface{}{"locked_secret": nil, "locked_at": nil})
	err := db.Error
	result := db.RowsAffected

	if err != nil {
		log.Warn("database error during the release of goods: ", err)
		http.Error(w, "the secret could not be validated", http.StatusInternalServerError)
		return
	}

	if result <= 0 {
		log.Warn("no goods found")
		http.Error(w, "there are no goods to release", http.StatusNotFound)
		return
	}

	log = log.WithField("count", result)

	lib.Write(w, map[string]int64{"count": result})
	log.Info("done")
}
