package create_tables

import (
	"github.com/wbrijesh/encompass/database"
	"github.com/wbrijesh/encompass/helpers"
)

func CreateWebsitesTable() (err error) {
	db, err := database.GetDatabase()
	_ = helpers.HandleError(err, helpers.Panic)

	_, err = db.Query(CreateWebsitesTableQuery)
	_ = helpers.HandleError(err, helpers.Panic)

	return nil
}

var CreateWebsitesTableQuery = `
CREATE TABLE IF NOT EXISTS "websites" (
  "id"         TEXT NOT NULL PRIMARY KEY,
  "name"       TEXT NOT NULL,
  "domain"     TEXT NOT NULL,
  "created_at" TEXT NOT NULL,
  "updated_at" TEXT,
  "userId"     TEXT,
  FOREIGN KEY ("userId") REFERENCES "User" ("id")
);
`
