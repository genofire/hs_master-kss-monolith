package runtime

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/worker"
	"github.com/genofire/hs_master-kss-monolith/models"
)

// create a worker to unlock goods which are locked by clients
func NewGoodReleaseWorker(grc models.GoodReleaseConfig) *worker.Worker {
	return worker.NewWorker(grc.Every.Duration, func() {
		goodRelease(grc.After.Duration)
	})
}

func goodRelease(unlockAfter time.Duration) int64 {
	res := database.Write.Model(&models.Good{}).Where("locked_secret is not NULL and locked_at < ?", time.Now().Add(-unlockAfter)).Updates(map[string]interface{}{"locked_secret": "", "locked_at": nil})
	return res.RowsAffected
}
