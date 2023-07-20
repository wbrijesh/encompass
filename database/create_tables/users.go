package create_tables

import (
	"github.com/wbrijesh/encompass/database"
	"github.com/wbrijesh/encompass/helpers"
)

var CreateUsersTableQuery string = `
CREATE TABLE IF NOT EXISTS "users" (
  "id"         TEXT NOT NULL PRIMARY KEY,
  "name"       TEXT NOT NULL,
  "email"      TEXT NOT NULL,
  "password"   TEXT NOT NULL,
  "created_at" TEXT NOT NULL,
  "updated_at" TEXT
);
`

func CreateUsersTable() (err error) {
	db, err := database.GetDatabase()
	_ = helpers.HandleError(err, helpers.Panic)

	_, err = db.Query(CreateUsersTableQuery)
	_ = helpers.HandleError(err, helpers.Panic)

	return nil
}
