package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
)

// this stock microservice manage goods
type Good struct {
	ID        int64
	ProductID int64
	Position  string
	Comment   string
	FouledAt  *time.Time

	RecievedAt *time.Time `sql:"default:current_timestamp"`
	// Make it temporary unusable
	LockedAt     *time.Time
	LockedSecret string `json:"-"`
	// Make it unusable
	DeletedAt *time.Time
	Sended    bool
}

// generate database select which filtered locked goods
func (g *Good) FilterAvailable(db *gorm.DB) *gorm.DB {
	return db.Model(g).Where("locked_secret == '' OR locked_secret is NULL")
}

// lock the good, on a way, that it could not be used by other users
func (g *Good) Lock(secret string) {
	now := time.Now()
	g.LockedSecret = secret
	g.LockedAt = &now
}

// is this good locked?
func (g *Good) IsLock() bool {
	return len(g.LockedSecret) > 0
}

// unlock the good, that it could be usered again
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
