package models

import "github.com/genofire/hs_master-kss-monolith/lib/database"

type Good struct {
	ID        int64
	ProductID int64
	Comment   string
}

func init() {
	database.AddModel(&Good{})
}
