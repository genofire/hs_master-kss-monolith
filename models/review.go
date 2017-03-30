package models

import "github.com/genofire/hs_master-kss-monolith/lib/database"

type Review struct {
	ID             int64
	ProductID      int64
	LocaleLanguage string
	FirstName      string
	LastName       string
	RatingStars    int64
	Text           string
}

func (r *Review) DisplayName() string {
	if len(r.FirstName) > 0 {
		if len(r.LastName) > 0 {
			last := []byte(r.LastName)
			return r.FirstName + " " + string(last[0]) + "."
		}
		return r.FirstName
	}
	return "Anonymous"
}

func init() {
	database.AddModel(&Review{})
}
