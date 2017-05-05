// Package with the mostly static content (models) of this microservice
package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
)

// Goods managed in this stock microservice
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

// Function to enerate a database and select locked goods with a filter
func (g *Good) FilterAvailable(db *gorm.DB) *gorm.DB {
	return db.Model(g).Where("locked_secret == '' OR locked_secret is NULL")
}

// Function to lock a good, so that it cannot be locked or bought by other users
func (g *Good) Lock(secret string) {
	now := time.Now()
	g.LockedSecret = secret
	g.LockedAt = &now
}

// Function to check if a good is locked
func (g *Good) IsLock() bool {
	return len(g.LockedSecret) > 0
}

// Function to unlock a good
func (g *Good) Unlock(secret string) error {
	if g.LockedSecret == secret {
		g.LockedSecret = ""
		g.LockedAt = nil
		return nil
	}
	return errors.New("wrong secret")
}

// Function to initialize the database
func init() {
	database.AddModel(&Good{})
}
