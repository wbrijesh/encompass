package create_tables

import (
	"github.com/wbrijesh/encompass/database"
	"github.com/wbrijesh/encompass/helpers"
)

var CreateSessionsTableQuery = `
CREATE TABLE IF NOT EXISTS "sessions" (
  "id"         TEXT NOT NULL PRIMARY KEY,
  "created_at" TEXT NOT NULL,
  "visitorId"  TEXT,
  FOREIGN KEY ("visitorId") REFERENCES "Visitor" ("id")
);
`

func CreateSessionsTable() (err error) {
	db, err := database.GetDatabase()
	_ = helpers.HandleError(err, helpers.Panic)

	_, err = db.Query(CreateSessionsTableQuery)
	_ = helpers.HandleError(err, helpers.Panic)

	return nil
}
