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
	Count     int64 `json:"count"`
}

func lockGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	secret := r.Header.Get("secret")
	log = log.WithField("lSecret", secret)

	tx := database.Write.Begin()
	// TODO the logic
	if tx.Error != nil {
		tx.Rollback()
		log.Warn("good not found")
		http.Error(w, "the good was not found in database", http.StatusNotFound)
		return
	}
	// TODO the logic

	lib.Write(w, map[string]int64{"count": 0})
	tx.Commit()
	log.Info("done")

}

func releaseGoods(w http.ResponseWriter, r *http.Request) {
	log := logger.HTTP(r)
	secret := r.Header.Get("secret")
	log = log.WithField("lSecret", secret)

	db := database.Write.Model(&models.Good{}).Where(&models.Good{LockedSecret: secret}).Updates(map[string]interface{}{"locked_secret": nil, "locked_at": nil})
	err := db.Error
	result := db.RowsAffected

	if err != nil {
		log.Warn("database error during release goods: ", err)
		http.Error(w, "secret could not validate", http.StatusInternalServerError)
		return
	}

	if result <= 0 {
		log.Warn("no goods found")
		http.Error(w, "no goods found to release", http.StatusNotFound)
		return
	}

	log = log.WithField("count", result)

	lib.Write(w, map[string]int64{"count": result})
	log.Info("done")
}
