package database

import (
	"database/sql"
	"fmt"
	"github.com/wbrijesh/encompass/helpers"
	"sync"
)

var once sync.Once
var db *sql.DB

func GetDatabase() (*sql.DB, error) {
	once.Do(func() {
		DbUrl, err := helpers.GetEnvVariable("DB_URL")
		_ = helpers.HandleError(err, helpers.Panic)

		conn, err := sql.Open("libsql", DbUrl)
		_ = helpers.HandleError(err, helpers.Panic)

		db = conn
	})

	if db == nil {
		return nil, fmt.Errorf("failed to initialize database")
	}

	return db, nil
}
