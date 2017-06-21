// Package with supporting functionality to run the microservice
package runtime

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// Function to automatically remove goods, if they are fouled
func GoodFouled() int {
	var goods []*models.Good
	var g models.Good
	g.FilterAvailable(database.Read).Where("fouled_at <= ?", time.Now()).Find(&goods)
	now := time.Now()

	for _, good := range goods {
		good.FouledDelete = true
		good.DeletedAt = &now
		database.Write.Save(&good)
	}
	return len(goods)
}
