package helpers

import "net/http"

func ServeHandler(path string, resolver func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		resolver(w, r)
	})
}
