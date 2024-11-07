package helper

import (
	"database/sql"

	"github.com/fajarherdian22/apiweb/repository"
)

func NullHandler(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func HandlingNullData(data []repository.DemarcationSiteLink) bool {
	return len(data) == 0
}
