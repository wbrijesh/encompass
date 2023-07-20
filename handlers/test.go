package handlers

import (
	"github.com/wbrijesh/encompass/helpers"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	_ = helpers.HandleError(err, helpers.Log)
}
