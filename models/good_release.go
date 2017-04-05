package models

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/worker"
)

type GoodReleaseConfig struct {
	After Duration `toml:"after"`
	Every Duration `toml:"every"`
}

func NewGoodReleaseWorker(grc GoodReleaseConfig) *worker.Worker {
	return worker.NewWorker(grc.Every.Duration, func() {
		goodRelease(grc.After.Duration)
	})
}

func goodRelease(unlockAfter time.Duration) int64 {
	res := database.Write.Model(&Good{}).Where("locked_secret is not NULL and locked_at < ?", time.Now().Add(-unlockAfter)).Updates(map[string]interface{}{"locked_secret": "", "locked_at": nil})
	return res.RowsAffected
}
