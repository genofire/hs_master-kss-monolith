package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
)

type Good struct {
	ID        int64
	ProductID int64
	Position  string
	Comment   string
	FouledAt  *time.Time

	RecievedAt *time.Time `sql:"default:current_timestamp"`
	// Make it temporary unusable
	LockedAt     *time.Time
	LockedSecret string
	// Make it unusable
	DeletedAt *time.Time
	Sended    bool
}

func (g *Good) FilterAvailable(db *gorm.DB) *gorm.DB {
	return db.Where("locked_secret is NULL deleted_at is NULL and send_at is NULL")
}

func (g *Good) Lock(secret string) {
	now := time.Now()
	g.LockedSecret = secret
	g.LockedAt = &now
}
func (g *Good) IsLock() bool {
	return len(g.LockedSecret) > 0
}
func (g *Good) Unlock(secret string) error {
	if g.LockedSecret == secret {
		g.LockedSecret = ""
		g.LockedAt = nil
		return nil
	}
	return errors.New("wrong secret")
}

func init() {
	database.AddModel(&Good{})
}
