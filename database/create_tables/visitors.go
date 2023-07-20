package create_tables

import (
	"github.com/wbrijesh/encompass/database"
	"github.com/wbrijesh/encompass/helpers"
)

var CreateVisitorsTableQuery = `
CREATE TABLE IF NOT EXISTS "visitors" (
  "id"               TEXT NOT NULL PRIMARY KEY,
  "browser"          TEXT,
  "operating_system" TEXT,
  "resolution"       TEXT,
  "device"           TEXT,
  "language"         TEXT,
  "country"          TEXT,
  "website_id"       TEXT,
  "created_at"       TEXT NOT NULL,
  "updated_at"       TEXT,
  FOREIGN KEY ("website_id") REFERENCES "Website" ("id")
);
`

func CreateVisitorsTable() (err error) {
	db, err := database.GetDatabase()
	_ = helpers.HandleError(err, helpers.Panic)

	_, err = db.Query(CreateVisitorsTableQuery)
	_ = helpers.HandleError(err, helpers.Panic)

	return nil
}
