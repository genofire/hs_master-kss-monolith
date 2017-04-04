package models

import (
	"time"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
)

type GoodReleaseConfig struct {
	After Duration `toml:"after"`
	Timer Duration `toml:"timer"`
}

type GoodReleaseWorker struct {
	unlockTimer time.Duration
	unlockAfter time.Duration
	quit        chan struct{}
}

func NewGoodReleaseWorker(grc GoodReleaseConfig) (rw *GoodReleaseWorker) {
	rw = &GoodReleaseWorker{
		unlockTimer: grc.Timer.Duration,
		unlockAfter: grc.After.Duration,
		quit:        make(chan struct{}),
	}
	return
}

func (rw *GoodReleaseWorker) Start() {
	ticker := time.NewTicker(rw.unlockTimer)
	for {
		select {
		case <-ticker.C:
			count := goodRelease(rw.unlockAfter)
			log.Log.WithField("count", count).Info("goods released")
		case <-rw.quit:
			ticker.Stop()
			return
		}
	}
}

func (rw *GoodReleaseWorker) Close() {
	close(rw.quit)
}

func goodRelease(unlockAfter time.Duration) int64 {
	res := database.Write.Model(&Good{}).Where("locked_secret is not NULL and locked_at < ?", time.Now().Add(-unlockAfter)).Updates(map[string]interface{}{"locked_secret": "", "locked_at": nil})
	return res.RowsAffected
}
