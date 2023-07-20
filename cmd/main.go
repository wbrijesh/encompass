package main

import (
	"fmt"
	_ "github.com/libsql/libsql-client-go/libsql"
	"github.com/wbrijesh/encompass/database/create_tables"
	"github.com/wbrijesh/encompass/handlers"
	"github.com/wbrijesh/encompass/helpers"
	"net/http"
)

func main() {
	ServerPort, err := helpers.GetEnvVariable("SERVER_PORT")
	err = create_tables.CreateAllTablesIfNotExists()
	_ = helpers.HandleError(err, helpers.Panic)

	helpers.ServeHandler("/", handlers.TestHandler)

	fmt.Println("server listening on port 4000")
	err = http.ListenAndServe(ServerPort, nil)
	_ = helpers.HandleError(err, helpers.Log)
}
